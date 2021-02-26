package config

import (
	"flag"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func ReadConf() {
	// 读取配置文件
	viper.SetConfigName("auth")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error while reading config file: " + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO when config file changes, what you expect to do in here
	})

	// 绑定环境变量
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取命令行参数
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic("Fatal error while bind pflags: " + err.Error())
	}
}

func Host() string         { return viper.GetString("server.host") }
func Port() int            { return viper.GetInt("server.port") }
func NeedInviteCode() bool { return viper.GetBool("server.need_invite_code") }
func DbConf() (string, string, int, string, string, string, int, int) {
	return viper.GetString("db.driver"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.lib"),
		viper.GetInt("db.max_idle_conn"),
		viper.GetInt("db.max_open_conn")
}
func DbLib() string {
	return viper.GetString("db.lib")
}
