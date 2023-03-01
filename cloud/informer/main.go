package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
)

func main() {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		klog.Fatal("read config error, config is empty")
		return
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)

	// 1 config
	// config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 2 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3 get informer
	// 获取一个namespace 一个获取全部 namespace
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("kube-system"))
	// factory := informers.NewSharedInformerFactory(clientset, 0) //resync 时间
	informer := factory.Core().V1().Pods().Informer()

	// 4 add event handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("update event")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete event")
		},
	})

	// 5 start informer
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
