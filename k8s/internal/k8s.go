package internal

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/liuqianhong6007/demo/k8s/config"
)

var gClientset *kubernetes.Clientset

func InitK8s() {
	var (
		err error
		cfg *rest.Config
	)
	if config.Cfg().Inner { // 在 kubernetes 内部使用
		cfg, err = rest.InClusterConfig()
	} else { // 在 kubernetes 外部使用
		cfg, err = clientcmd.BuildConfigFromFlags("", config.Cfg().CfgPath)
	}
	if err != nil {
		panic(err)
	}
	gClientset, err = kubernetes.NewForConfig(cfg)
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
