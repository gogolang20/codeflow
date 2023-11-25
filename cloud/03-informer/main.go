package main

import (
	"fmt"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

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
	// 3.1 获取指定 namespace
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))
	
	// 3.2 获取全部 namespace
	// 第二个参数: resync 时间
	// factory := informers.NewSharedInformerFactory(clientset, 0)

	informer := factory.Core().V1().Pods().Informer()

	// 4 add event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update event")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete event")
		},
	})

	// 5 start informer
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
