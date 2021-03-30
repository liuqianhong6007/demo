package internal

import (
	"github.com/liuqianhong6007/demo/k8s/com"
	"k8s.io/client-go/kubernetes"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Inner   bool   `yaml:"inner"`
	CfgPath string `yaml:"cfg_path"`
}

var gConfig Config

func ReadConfig(filepath string) {
	com.ReadConfig(filepath, &gConfig)
}

func Cfg() *Config {
	return &gConfig
}

var gClientset *kubernetes.Clientset

func InitK8sClientset(inner bool, cfgPath string) {
	gClientset = com.NewKubernetesClientset(inner, cfgPath)
}

func K8sClientset() *kubernetes.Clientset {
	return gClientset
}
