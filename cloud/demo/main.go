package main

import (
	"context"
	"fmt"
	"istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/describe"
	"log"
)

var (
	clientset       *kubernetes.Clientset
	dynamicClient   dynamic.Interface
	discoveryClient *discovery.DiscoveryClient
	istioClient     *versioned.Clientset

	describes describe.PodDescriber
)

const (
	namespace = "default"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln("get config error: ", err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("create clientset error: ", err)
	}

	dynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalln("create dynamicClient error: ", err)
	}

	discoveryClient, err = discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		log.Fatalln("create dynamicClient error: ", err)
	}

	describes = describe.PodDescriber{
		Interface: clientset,
	}

	istioClient, err = versioned.NewForConfig(config)
	if err != nil {
		log.Fatalln("create istioClient error: ", err)
	}

	// TODO istio client
	list, _ := istioClient.NetworkingV1alpha3().VirtualServices(namespace).List(context.Background(), metav1.ListOptions{})
	fmt.Println(list)

	// TODO describes pod
	str, _ := describes.Describe(namespace, "busybox", describe.DescriberSettings{})
	fmt.Println(str)

	// TODO CRD
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1alpha3",
		Resource: "virtualservices",
	}
	unstructuredList, _ := dynamicClient.Resource(virtualServiceGVR).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	fmt.Println(unstructuredList)
}
