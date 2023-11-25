package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		klog.Fatal(err)
	}

	// client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	// get data
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "busybox", v1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	} else {
		fmt.Println("clientset: ", pod.Name)
	}
}
