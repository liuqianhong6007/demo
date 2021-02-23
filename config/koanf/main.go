package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/config/koanf/config"
)

func main() {
	os.Setenv("GAME_SERVICE.SERVER_ID", "20")

	config.SetEnvPrefix("")
	config.ReadConf("demo", "yml")

	fmt.Println("ServiceID: ", config.Get("game_service.server_id"))
	fmt.Println("ServicePort: ", config.Get("game_service.port"))
	fmt.Println("PassportKey: ", config.Get("game_service.passport_key"))
	fmt.Println("GameplayManagerAddr: ", config.Get("game_service.gameplay_manager_addr"))
	fmt.Println("EtcdEndpoint: ", config.Get("game_service.etcd_endpoint"))

	// unmarshal 支持将值映射到一个结构体
	var conf Config
	config.Unmarshal("", &conf)
	fmt.Println(conf)
}

type Config struct {
	GameService GameService `koanf:"game_service"`
	Fastway     Fastway     `koanf:"fastway"`
	DB          DB          `koanf:"db"`
	Log         Log         `koanf:"log"`
}

type GameService struct {
	ServiceID           int    `koanf:"server_id"`
	Port                int    `koanf:"port"`
	PassportKey         string `koanf:"passport_key"`
	GameplayManagerAddr string `koanf:"gameplay_manager_addr"`
}

type Fastway struct {
	ServerAddr string `koanf:"server_addr"`
	AuthKey    string `koanf:"auth_key"`
}

type DB struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
}

type Log struct {
	Dir         string `koanf:"dir"`
	Name        string `koanf:"name"`
	MaxSize     int    `koanf:"max_size"`
	MaxAge      int    `koanf:"max_age"`
	MaxBackups  int    `koanf:"max_backups"`
	Compress    bool   `koanf:"compress"`
	LogChanSize int    `koanf:"log_chan_size"`
	ReleaseMode bool   `koanf:"release_mode"`
}
