package main

import (
	"log"

	"codeflow/cloud/05-operator/pkg"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

/*
一个小项目，关于 operator
*/
func main() {
	// 1 config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		// controller 运行在集群内部，找不到config文件，进入报错
		// 使用集群内部对象创建
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("get config error: ", err)
		}
		config = inClusterConfig
	}

	// 2 client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("create clientset error: ", err)
	}

	// 3 informer
	factory := informers.NewSharedInformerFactory(clientset, 0)
	// 创建 factory 关注的资源对象 informer
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()

	// 4 add event handler
	// 抽离到 controller 中
	controller := pkg.NewController(clientset, serviceInformer, ingressInformer)

	// 5 informer start
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	// 启动controller
	controller.Run(stopCh)
}
