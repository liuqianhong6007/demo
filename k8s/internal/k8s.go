package internal

import (
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var gClientset *kubernetes.Clientset

func InitK8s(cfgPath string) {
	config, err := clientcmd.BuildConfigFromFlags("", cfgPath)
	if err != nil {
		panic(err)
	}
	gClientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}

func Clientset() *kubernetes.Clientset {
	if gClientset == nil {
		panic("uninitialized kubernetes client set")
	}
	return gClientset
}

func CoreV1() v1.CoreV1Interface {
	return Clientset().CoreV1()
}
