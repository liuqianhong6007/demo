package config

import (
	"github.com/spf13/viper"
)

var (
	kServiceId           = "game_service.server_id"
	kServicePort         = "game_service.port"
	kPassportKey         = "game_service.passport_key"
	kGameplayManagerAddr = "game_service.gameplay_manager_addr"
	kEtcdEndpoint        = "game_service.etcd_endpoint"

	kFastwayAddr    = "fastway.server_addr"
	kFastwayAuthkey = "fastway.auth_key"

	kDBHost     = "db.host"
	kDBPort     = "db.port"
	kDBUser     = "db.user"
	kDBPassword = "db.password"
	kDBName     = "db.name"

	kLogDir         = "log.dir"
	kLogName        = "log.name"
	kLogMaxSize     = "log.max_size"
	kLogMaxAge      = "log.max_age"
	kLogMaxBackups  = "log.max_backups"
	kLogCompress    = "log.compress"
	kLogChanSize    = "log.log_chan_size"
	kLogReleaseMode = "log.release_mode"
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
