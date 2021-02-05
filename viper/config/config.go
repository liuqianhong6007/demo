package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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

func setDefault() {
	viper.SetDefault(kServiceId, 1)
	viper.SetDefault(kServicePort, 8000)
	viper.SetDefault(kEtcdEndpoint, "http://127.0.0.1:2379")
}

func readInConfig() {
	viper.SetConfigName("game_service")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while reading config file: %s\n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO when config file changes, what you expect to do in here
	})
}

func setOverride() {
	viper.Set(kLogDir, "log")
}

func readEnv() {
	viper.AutomaticEnv()
}

func readFlag() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while bind pflags: %s\n", err))
	}
}

func readEtcd(path string) {
	err := viper.AddRemoteProvider("etcd", viper.GetString(kEtcdEndpoint), path)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while adding remote provider: %s\n", err))
	}
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while reading remote config: %s\n", err))
	}
	err = viper.WatchRemoteConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while watch remote config: %s\n", err))
	}
}

func ReadConf() {
	readFlag()
	readEnv()
	readInConfig()
	setDefault()
	setOverride()
	readEtcd("/etc/conf/")
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
