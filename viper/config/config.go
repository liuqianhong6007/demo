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

var (
	setDefaultMap  = make(map[string]interface{})
	setOverrideMap = make(map[string]interface{})
)

func CombineKey(prefix string, paths ...string) string {
	pathArray := []string{prefix}
	for _, path := range paths {
		pathArray = append(pathArray, path)
	}
	return strings.Join(pathArray, ".")
}

func SetDefault(key string, value interface{}) {
	if _, ok := setDefaultMap[key]; ok {
		panic("duplicate set default key: " + key)
	}
	setDefaultMap[key] = value
}

func SetOverride(key string, value interface{}) {
	if _, ok := setOverrideMap[key]; ok {
		panic("duplicate set override key: " + key)
	}
	setOverrideMap[key] = value
}

func ReadConf(configName, configTyp string) {
	// read flag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while bind pflags: %s\n", err))
	}

	// read env
	viper.AutomaticEnv()

	// read config file
	viper.SetConfigName(configName)
	viper.SetConfigType(configTyp)
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
	for k, v := range setDefaultMap {
		viper.SetDefault(k, v)
	}

	// set override
	for k, v := range setOverrideMap {
		viper.Set(k, v)
	}
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
