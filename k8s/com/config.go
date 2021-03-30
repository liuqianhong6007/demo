package com

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func ReadConfig(filepath string, cfg interface{}) {
	if reflect.TypeOf(cfg).Kind() != reflect.Ptr {
		panic("cfg must be a point")
	}
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
	err = viper.Unmarshal(cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})
	if err != nil {
		panic("Fatal error while unmarshal config")
	}
}
