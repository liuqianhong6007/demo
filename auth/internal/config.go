package internal

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
func Secret() string       { return viper.GetString("server.secret") }

func DbDriver() string            { return viper.GetString("db.driver") }
func DbHost() string              { return viper.GetString("db.host") }
func DbPort() int                 { return viper.GetInt("db.port") }
func DbUser() string              { return viper.GetString("db.user") }
func DbPassword() string          { return viper.GetString("db.password") }
func DbLib() string               { return viper.GetString("db.lib") }
func DbMaxIdleConn() int          { return viper.GetInt("db.max_idle_conn") }
func DbMaxOpenConn() int          { return viper.GetInt("db.max_open_conn") }
func CastbinModelPath() string    { return viper.GetString("castbin.model_path") }
func CastbinPolicyDriver() string { return viper.GetString("castbin.policy_driver") }
func CastbinPolicyPath() string   { return viper.GetString("castbin.policy_path") }
