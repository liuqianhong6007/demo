package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func InitK8s(cfgPath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", cfgPath)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientSet
}
