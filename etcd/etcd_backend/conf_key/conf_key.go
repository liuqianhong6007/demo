package conf_key

import (
	"github.com/liuqianhong6007/demo/viper/config"
	"github.com/spf13/viper"
)

var (
	kHost     = config.CombineKey("host")
	kPort     = config.CombineKey("port")
	kEtcdAddr = config.CombineKey("etcd_addr")
)

func init() {
	config.SetDefault(kHost, "0.0.0.0")
	config.SetDefault(kPort, "8101")
	config.SetDefault(kEtcdAddr, "127.0.0.1:2379")
	config.ReadConf("etcd_backend", "yaml")
}

func Host() string {
	return viper.GetString(kHost)
}

func Port() int {
	return viper.GetInt(kPort)
}

func EtcdAddr() string {
	return viper.GetString(kEtcdAddr)
}
