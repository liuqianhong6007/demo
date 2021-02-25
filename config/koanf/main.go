package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/config/koanf/config"
)

func main() {
	os.Setenv("GAMESERVICE_SERVERID", "20")

	config.SetEnvPrefix("")
	config.ReadConf("demo", "yml")

	// unmarshal 支持将值映射到一个结构体
	var conf Config
	config.Unmarshal("", &conf)
	buf, _ := json.Marshal(conf)
	fmt.Println(string(buf))
}

type Config struct {
	GameService GameService `koanf:"gameservice"`
	Fastway     Fastway     `koanf:"fastway"`
	DB          DB          `koanf:"db"`
	Log         Log         `koanf:"log"`
}

type GameService struct {
	ServiceID           int    `koanf:"serverid"`
	Port                int    `koanf:"port"`
	PassportKey         string `koanf:"passportkey"`
	GameplayManagerAddr string `koanf:"gameplaymanageraddr"`
}

type Fastway struct {
	ServerAddr string `koanf:"serveraddr"`
	AuthKey    string `koanf:"authkey"`
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
	MaxSize     int    `koanf:"maxsize"`
	MaxAge      int    `koanf:"maxgge"`
	MaxBackups  int    `koanf:"maxbackups"`
	Compress    bool   `koanf:"compress"`
	LogChanSize int    `koanf:"logchansize"`
	ReleaseMode bool   `koanf:"releasemode"`
}
