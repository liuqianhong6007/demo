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
	kServiceId           = combineKey("game_service", "server_id")
	kServicePort         = combineKey("game_service", "port")
	kPassportKey         = combineKey("game_service", "passport_key")
	kGameplayManagerAddr = combineKey("game_service", "gameplay_manager_addr")
	kEtcdEndpoint        = combineKey("game_service", "etcd_endpoint")
	kFastwayAddr         = combineKey("fastway", "server_addr")
	kFastwayAuthkey      = combineKey("fastway", "auth_key")
	kDBHost              = combineKey("db", "host")
	kDBPort              = combineKey("db", "port")
	kDBUser              = combineKey("db", "user")
	kDBPassword          = combineKey("db", "password")
	kDBName              = combineKey("db", "name")
	kLogDir              = combineKey("log", "dir")
	kLogName             = combineKey("log", "name")
	kLogMaxSize          = combineKey("log", "max_size")
	kLogMaxAge           = combineKey("log", "max_age")
	kLogMaxBackups       = combineKey("log", "max_backups")
	kLogCompress         = combineKey("log", "compress")
	kLogChanSize         = combineKey("log", "log_chan_size")
	kLogReleaseMode      = combineKey("log", "release_mode")
)

func combineKey(prefix string, paths ...string) string {
	pathArray := []string{prefix}
	for _, path := range paths {
		pathArray = append(pathArray, path)
	}
	return strings.Join(pathArray, ".")
}

func ReadConf() {
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
	viper.SetConfigName("game_service")
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
	viper.SetDefault(kServicePort, 8000)
	viper.SetDefault(kEtcdEndpoint, "http://127.0.0.1:2379")

	// set override
	viper.Set(kLogDir, "log")

	// read from remote
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(viper.GetString(kEtcdEndpoint), ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error while new etcd client: %s\n", err))
	}

	ctx, cancleFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancleFunc()

	rsp, err := client.Get(ctx, "/conf/game_service")
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

func ServiceId() string {
	return viper.GetString(kServiceId)
}

func ServicePort() int {
	return viper.GetInt(kServicePort)
}

func PassportKey() string {
	return viper.GetString(kPassportKey)
}

func GameplayManagerAddr() string {
	return viper.GetString(kGameplayManagerAddr)
}

func EtcdEndpoint() string {
	return viper.GetString(kEtcdEndpoint)
}

func FastwayAddr() string {
	return viper.GetString(kFastwayAddr)
}

func FastwayAuthkey() string {
	return viper.GetString(kFastwayAuthkey)
}

func DBHost() string {
	return viper.GetString(kDBHost)
}

func DBPort() int {
	return viper.GetInt(kDBPort)
}

func DBUser() string {
	return viper.GetString(kDBUser)
}

func DBPassword() string {
	return viper.GetString(kDBPassword)
}

func DBName() string {
	return viper.GetString(kDBName)
}

func LogDir() string {
	return viper.GetString(kLogDir)
}

func LogName() string {
	return viper.GetString(kLogName)
}

func LogMaxSize() int {
	return viper.GetInt(kLogMaxSize)
}

func LogMaxAge() int {
	return viper.GetInt(kLogMaxAge)
}

func LogMaxBackups() int {
	return viper.GetInt(kLogMaxBackups)
}

func LogCompress() bool {
	return viper.GetBool(kLogCompress)
}

func LogChanSize() int {
	return viper.GetInt(kLogChanSize)
}

func LogReleaseMode() bool {
	return viper.GetBool(kLogReleaseMode)
}
