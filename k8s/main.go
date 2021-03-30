package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/liuqianhong6007/demo/k8s/api"
	"github.com/liuqianhong6007/demo/k8s/config"
	"github.com/liuqianhong6007/demo/k8s/internal"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "./config.yaml", "config file path")
	flag.Parse()
}

func main() {
	config.ReadConfig(cfgPath)

	// init k8s client
	internal.InitK8s()

	r := gin.Default()
	internal.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", config.Cfg().Host, config.Cfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
