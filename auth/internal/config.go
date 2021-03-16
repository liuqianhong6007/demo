package internal

import (
	"github.com/liuqianhong6007/viper_x"
)

type Config struct {
	Server ServerCfg `viper:"server"`
	Db     DbCfg     `viper:"db"`
	Casbin CasbinCfg `viper:"casbin"`
}

type ServerCfg struct {
	Host           string `viper:"host"`
	Port           int    `viper:"port"`
	NeedInviteCode bool   `viper:"need_invite_code"`
	Secret         string `viper:"secret"`
}

type DbCfg struct {
	Driver      string `viper:"driver"`
	Host        string `viper:"host"`
	Port        int    `viper:"port"`
	User        string `viper:"user"`
	Password    string `viper:"password"`
	Lib         string `viper:"lib"`
	MaxIdleConn int    `viper:"max_idle_conn"`
	MaxOpenConn int    `viper:"max_open_conn"`
}

type CasbinCfg struct {
	ModelPath    string `viper:"model_path"`
	PolicyDriver string `viper:"policy_driver"`
	PolicyPath   string `viper:"policy_path"`
}

var gConfig Config

func ReadConf() {
	viper_x.ReadConf("auth", &gConfig)
}

func GetConfig() *Config {
	return &gConfig
}
