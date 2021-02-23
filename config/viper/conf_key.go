package main

import (
	"github.com/spf13/viper"

	"github.com/liuqianhong6007/demo/config/viper/config"
)

var (
	kServiceId           = config.CombineKey("game_service", "server_id")
	kServicePort         = config.CombineKey("game_service", "port")
	kPassportKey         = config.CombineKey("game_service", "passport_key")
	kGameplayManagerAddr = config.CombineKey("game_service", "gameplay_manager_addr")
	kEtcdEndpoint        = config.CombineKey("game_service", "etcd_endpoint")
	kFastwayAddr         = config.CombineKey("fastway", "server_addr")
	kFastwayAuthkey      = config.CombineKey("fastway", "auth_key")
	kDBHost              = config.CombineKey("db", "host")
	kDBPort              = config.CombineKey("db", "port")
	kDBUser              = config.CombineKey("db", "user")
	kDBPassword          = config.CombineKey("db", "password")
	kDBName              = config.CombineKey("db", "name")
	kLogDir              = config.CombineKey("log", "dir")
	kLogName             = config.CombineKey("log", "name")
	kLogMaxSize          = config.CombineKey("log", "max_size")
	kLogMaxAge           = config.CombineKey("log", "max_age")
	kLogMaxBackups       = config.CombineKey("log", "max_backups")
	kLogCompress         = config.CombineKey("log", "compress")
	kLogChanSize         = config.CombineKey("log", "log_chan_size")
	kLogReleaseMode      = config.CombineKey("log", "release_mode")
)

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
