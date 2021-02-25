package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gopkg.in/yaml.v2"
)

func ReadConf() {
	// read flag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while bind pflags: %s\n", err))
	}

	// read env
	// viper.AutomaticEnv()
	viper.BindEnv(kServiceId, "GAME_SERVICE_SERVER_ID")
	viper.BindEnv(kServicePort, "GAME_SERVICE_PORT")

	// read config file
	viper.SetConfigName("demo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while reading config file: %s\n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO when config file changes, what you expect to do in here
	})

	// set default
	viper.SetDefault(kServiceId, 1)
	//viper.SetDefault(kServicePort, 8000)
	viper.SetDefault(kEtcdEndpoint, "http://127.0.0.1:2379")

	// set override
	viper.Set(kLogDir, "log")
}

func ReadFromEtcd(endpoints, path string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error while new etcd client: %s\n", err))
	}

	ctx, cancleFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancleFunc()

	rsp, err := client.Get(ctx, path)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while get etcd key value: %s\n", err))
	}
	var val interface{}
	err = yaml.Unmarshal(rsp.Kvs[0].Value, &val)
	if err != nil {
		panic(fmt.Sprintf("Faral error while etcd value unmarshal: %s\n", err))
	}
	log.Println("Etcd value: ", val)
}
