package config

import (
	"flag"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	kHost     = "host"
	kPort     = "port"
	kEtcdAddr = "etcd_addr"
)

func ReadConf() {
	// 设置默认值
	viper.SetDefault(kHost, "0.0.0.0")
	viper.SetDefault(kPort, "8101")
	viper.SetDefault(kEtcdAddr, "127.0.0.1:2379")

	// 读取配置文件
	viper.SetConfigName("etcd_backend")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error while reading config file: " + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO when config file changes, what you expect to do in here
	})

	// 绑定环境变量
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取命令行参数
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic("Fatal error while bind pflags: " + err.Error())
	}
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
