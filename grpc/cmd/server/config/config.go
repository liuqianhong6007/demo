package config

import (
	"flag"
	"fmt"
)

type Config struct {
	ip   string
	port int
}

var gConfig *Config

func init() {
	gConfig = &Config{}
	flag.Parse()
	flag.StringVar(&gConfig.ip, "ip", "0.0.0.0", "grpc ip")
	flag.IntVar(&gConfig.port, "port", 8080, "grpc port")
}

func GrpcAddr() string {
	return fmt.Sprintf("%s:%d", gConfig.ip, gConfig.port)
}
