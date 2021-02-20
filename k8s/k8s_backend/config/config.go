package config

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	K8s    K8s    `yaml:"k8s"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type K8s struct {
	CfgPath string `yaml:"cfg_path"`
}

var (
	gConfig Config
)

func init() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "config path")
	flag.Parse()

	buff, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(buff, &gConfig)
	if err != nil {
		panic(err)
	}
}

func ServerCfg() Server {
	return gConfig.Server
}

func K8sCfg() K8s {
	return gConfig.K8s
}
