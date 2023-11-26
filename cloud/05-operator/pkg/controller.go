package pkg

import (
	"context"
	"reflect"
	"time"

	v14 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informer "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	workNum  = 5  // 开启goroutine 数量
	maxRetry = 10 // 最大重试次数
)

type controller struct {
	client        kubernetes.Interface            // 和kube-apiserver交互
	ingressLister v1.IngressLister                // 控制器操作的资源对象
	serviceLister coreLister.ServiceLister        // 控制器操作的资源对象
	queue         workqueue.RateLimitingInterface // 限速队列
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	// TODO 比较 annotation
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}

func (c *controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c *controller) enqueue(obj interface{}) {
	// cache.MetaNamespaceKeyFunc() 默认获取key的方法
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}

	c.queue.Add(key)
}

func (c *controller) deleteIngress(obj interface{}) {
	// 删除ingress时，如果service的annotation是存在的，并且是期望的
	// 那么ingress应该被重建
	ingress := obj.(*v12.Ingress)

	// 获取ingress的service
	ownerReference := v13.GetControllerOf(ingress)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind != "Service" {
		return
	}

	c.queue.Add(ingress.Namespace + "/" + ingress.Name)
}

// 从 queue 中取出数据消费
func (c *controller) Run(stopCh chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.worker, time.Minute, stopCh)
	}
	<-stopCh
}

// 不停的从workqueue中获取key，并处理
func (c *controller) worker() {
	for c.processNextItem() {
	}
}

// 从key中，将 obj对象 获取处理
func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item) // 处理完毕后移除 item

	key := item.(string)
	err := c.syncService(key)
	if err != nil {
		c.handlerError(key, err)
	}
	return true
}

// 协调资源状态
func (c *controller) syncService(key string) error {
	namespaceKey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	// 删除
	service, err := c.serviceLister.Services(namespaceKey).Get(name)
	if errors.IsNotFound(err) { // 表示已经删除，就不用管了
		return nil
	}
	if err != nil {
		return err
	}

	// 新增和删除
	_, ok := service.GetAnnotations()["ingress/http"]

	ingress, err := c.ingressLister.Ingresses(namespaceKey).Get(name)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	// 有 annotation对应的service 并且 没有ingress
	if ok && errors.IsNotFound(err) {
		// create ingress
		ig := c.constructIngress(service)
		// 通过client 创建
		_, err := c.client.NetworkingV1().Ingresses(namespaceKey).Create(context.Background(), ig, v13.CreateOptions{})
		return err
	} else if !ok && ingress != nil { // 没有 annotation对应的service 并且 有ingress
		// delete ingress
		return c.client.NetworkingV1().Ingresses(namespaceKey).Delete(context.Background(), name, v13.DeleteOptions{})
	}
	return nil
}

// 将key重新放入queue
func (c *controller) handlerError(key string, err error) {
	// 进行重试
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

// 构建一个 ingress
func (c *controller) constructIngress(service *v14.Service) *v12.Ingress {
	ingress := v12.Ingress{}

	// 需要指定 OwnerReference
	// https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/owners-dependents/
	ingress.ObjectMeta.OwnerReferences = []v13.OwnerReference{
		// ingress 的owner 是 service
		*v13.NewControllerRef(service, v14.SchemeGroupVersion.WithKind("Service")),
	}

	ingress.Name = service.Name
	ingress.Namespace = service.Namespace
	pathType := v12.PathTypePrefix
	inc := "nginx" // 需要添加
	ingress.Spec = v12.IngressSpec{
		IngressClassName: &inc,
		Rules: []v12.IngressRule{
			{
				Host: "",
				IngressRuleValue: v12.IngressRuleValue{
					HTTP: &v12.HTTPIngressRuleValue{
						Paths: []v12.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &pathType,
								Backend: v12.IngressBackend{
									Service: &v12.IngressServiceBackend{
										Name: service.Name,
										Port: v12.ServiceBackendPort{
											Number: 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return &ingress
}

func NewController(client kubernetes.Interface, serviceinformer informer.ServiceInformer, ingressinformer netInformer.IngressInformer) controller {
	c := controller{
		client:        client,
		ingressLister: ingressinformer.Lister(),
		serviceLister: serviceinformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressHandler"),
	}
	// 定义 event 处理方法
	serviceinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})
	ingressinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})

	return c
}
