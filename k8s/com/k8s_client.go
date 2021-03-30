package com

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubernetesClientset(inner bool, cfgPath string) *kubernetes.Clientset {
	var (
		err error
		cfg *rest.Config
	)
	if inner { // 在 kubernetes 内部使用
		cfg, err = rest.InClusterConfig()
	} else { // 在 kubernetes 外部使用
		cfg, err = clientcmd.BuildConfigFromFlags("", cfgPath)
	}
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	return clientset
}
