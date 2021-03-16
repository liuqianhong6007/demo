package config

import (
	"github.com/liuqianhong6007/viper_x"
)

type Config struct {
	Host     string `viper:"host"`
	Port     int    `viper:"port"`
	EtcdAddr string `viper:"etcd_addr"`
}

var gConfig Config

func ReadConf() {
	viper_x.ReadConf("etcd_backend", &gConfig)
}

func GetConfig() *Config {
	return &gConfig
}
