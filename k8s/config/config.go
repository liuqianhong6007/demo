package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var gConfig Config

func Cfg() *Config {
	return &gConfig
}

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Inner   bool   `yaml:"inner"`
	CfgPath string `yaml:"cfg_path"`
}

func ReadConfig(filepath string) {
	// read flag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while bind pflags: %s\n", err))
	}

	// read env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// read config file
	viper.SetConfigFile(filepath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error while reading config file: %s\n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {})
	err = viper.Unmarshal(&gConfig, func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})
	if err != nil {
		panic("Fatal error while unmarshal config")
	}
}
