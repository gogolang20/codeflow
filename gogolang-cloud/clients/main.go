// package developerportal
package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	"istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/kubectl/pkg/describe"
	"path/filepath"
)

var (
	kubeconfig      *string
	clientset       *kubernetes.Clientset
	dynamicClient   dynamic.Interface
	describes       describe.PodDescriber
	discoveryClient *discovery.DiscoveryClient
	istioClient     *versioned.Clientset
)

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "../testdata/gaia/testkubeconfig.yml", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logrus.Error("get config error: ", err)
		return
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error("create clientset error: ", err)
		return
	}

	dynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		logrus.Error("create dynamicClient error: ", err)
		return
	}

	describes = describe.PodDescriber{
		Interface: clientset,
	}

	discoveryClient, err = discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		logrus.Error("create dynamicClient error: ", err)
		return
	}

	istioClient, err = versioned.NewForConfig(config)
	if err != nil {
		logrus.Error("create istioClient error: ", err)
		return
	}
}

const (
	namespace   = "istio-system"
	podName     = ""
	serviceName = "prometheus"
	ingressName = ""
	crdName     = ""
)

func GetRequestauthentications() {
	vsList, err := istioClient.NetworkingV1alpha3().VirtualServices(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("Failed to get VirtualService in [%s] namespace: %s", namespace, err)
		return
	}

	for i, vs := range vsList.Items {
		logrus.Printf("Index: %d VirtualService Hosts: %+v [%s]\n", i, vs.Spec.GetHosts(), vs.Name)
	}

	resources, _, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return
	}

	for index, resource := range resources {
		if resource.Name == "security.istio.io" {
			fmt.Printf("resources name [%s], [%d]th\n", resource.Name, index)
		}
	}

}

func GetPod(namespace, name string) *v1.Pod {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		logrus.Error("[GetPod] error: ", err)
		return nil
	}

	for i, container := range pod.Spec.Containers {
		fmt.Printf("[GetPod] [%d]th container name is [%s]: ", i, container.Name)
		for index, env := range container.Env {
			fmt.Printf("The [%d]th key [%s] value is [%s].\n", index, env.Name, env.Value)
		}
	}

	return pod
}

func GetPodDescribe(namespace, name string) {
	des, err := describes.Describe(namespace, name, describe.DescriberSettings{})
	if err != nil {
		logrus.Error("[GetPodDescribe] error: ", err)
		return
	}

	fmt.Println(des)
}

func GetNamespace(name string) *v1.Namespace {
	ns, err := clientset.CoreV1().Namespaces().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		logrus.Error("[GetNamespace] error: ", err)
		return nil
	}

	return ns
}

func ListPods(namespace string) *v1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logrus.Error("[ListPods] error: ", err)
		return nil
	}

	return pods
}

func GetService(namespace, name string) *v1.Service {
	service, err := clientset.CoreV1().Services(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		logrus.Error("[GetService] error: ", err)
		return nil
	}

	return service
}

func ListService(namespace string) *v1.ServiceList {
	services, err := clientset.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		logrus.Error("[ListService] error: ", err)
		return nil
	}

	return services
}

/*
	CRD
*/

func ListVirtualService() []unstructured.Unstructured {
	// Create a GVR which represents an Istio Virtual Service.
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1alpha3",
		Resource: "virtualservices",
	}

	//  List all of the Virtual Services.
	virtualServices, err := dynamicClient.Resource(virtualServiceGVR).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil
	}

	return virtualServices.Items
}

func GetVirtualService(name string) {
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    "networking.istio.io",
		Version:  "v1alpha3",
		Resource: "virtualservices",
	}

	//  List all of the Virtual Services.
	virtualService, err := dynamicClient.Resource(virtualServiceGVR).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return
	}

	fmt.Println(virtualService.GetAPIVersion())
}
