package main

import (
	"fmt"
	"github.com/liuqianhong6007/demo/config/viper/config"
)

func main() {

	var conf Config

	config.ReadConf(&conf)

	fmt.Println(conf)
}

type Config struct {
	GameService GameService `viper:"game_service"`
	Fastway     Fastway     `viper:"fastway"`
	DB          DB          `viper:"db"`
	Log         Log         `viper:"log"`
}

type GameService struct {
	ServiceID           int    `viper:"server_id"`
	Port                int    `viper:"port"`
	PassportKey         string `viper:"passport_key"`
	GameplayManagerAddr string `viper:"gameplay_manager_addr"`
}

type Fastway struct {
	ServerAddr string `viper:"server_addr"`
	AuthKey    string `viper:"auth_key"`
}

type DB struct {
	Host     string `viper:"host"`
	Port     int    `viper:"port"`
	User     string `viper:"user"`
	Password string `viper:"password"`
	Name     string `viper:"name"`
}

type Log struct {
	Dir         string `viper:"dir"`
	Name        string `viper:"name"`
	MaxSize     int    `viper:"max_size"`
	MaxAge      int    `viper:"max_gge"`
	MaxBackups  int    `viper:"max_backups"`
	Compress    bool   `viper:"compress"`
	LogChanSize int    `viper:"log_chan_size"`
	ReleaseMode bool   `viper:"release_mode"`
}
