package config

import (
	"flag"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/spf13/pflag"
)

var (
	k         = koanf.New(".")
	envPrefix = ""
)

func SetEnvPrefix(prefix string) {
	if prefix != "" {
		envPrefix = prefix
	}
}

func ReadConf(configName, configTyp string) {
	// read config file
	filename := configName + "." + configTyp
	switch configTyp {
	case "json":
		if err := k.Load(file.Provider(filename), json.Parser()); err != nil {
			panic("error loading json config: " + err.Error())
		}
	case "yml", "yaml":
		if err := k.Load(file.Provider(filename), yaml.Parser()); err != nil {
			panic("error loading yaml config: " + err.Error())
		}
	case "toml":
		if err := k.Load(file.Provider(filename), toml.Parser()); err != nil {
			panic("error loading toml config: " + err.Error())
		}
	case "env":
		if err := k.Load(file.Provider(filename), dotenv.Parser()); err != nil {
			panic("error loading .env config: " + err.Error())
		}
	default:
		panic("not supported config type: " + configTyp)
	}

	// read from env
	if err := k.Load(env.Provider(envPrefix, "_", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, envPrefix))
	}), nil); err != nil {
		panic("error loading environment: " + err.Error())
	}

	// read from commandline
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := k.Load(posflag.Provider(pflag.CommandLine, ".", k), nil); err != nil {
		panic("error loading commandline: " + err.Error())
	}
}

func Get(key string) interface{} {
	return k.Get(key)
}

func Unmarshal(path string, o interface{}) {
	if err := k.Unmarshal(path, o); err != nil {
		panic("error unmarshalling: " + err.Error())
	}
}
