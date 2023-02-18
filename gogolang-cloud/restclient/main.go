package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		klog.Fatal("get config err:", err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	restclient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	var pod = v1.Pod{}
	err = restclient.Get().Namespace("default").Resource("pods").Name("test").Do(context.Background()).Into(&pod)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pod.Name)
	}
}
