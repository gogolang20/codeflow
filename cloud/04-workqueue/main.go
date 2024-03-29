package main

import (
	"fmt"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

/*
使用队列原因
	event 产生速度比 event 消费的速度快

1 通用队列
2 延时队列
3 限速队列

*/

func main() {
	// var kubeConfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "")
	// } else {
	// 	klog.Fatal("read config error, config is empty")
	// 	return
	// }
	// flag.Parse()
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)

	// 1 config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	// 2 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3 get informer
	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Core().V1().Pods().Informer()

	// 传入默认限速器 并给限速器命名
	rateLimitingQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller")

	// 4 add event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Println("[AddFunc] get key err: ", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update event")
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				fmt.Println("[UpdateFunc] get key err: ", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				fmt.Println("[DeleteFunc] get key err: ", err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
	})

	// 5 start informer
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
