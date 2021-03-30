package main

import (
	"flag"
	"fmt"
	"github.com/liuqianhong6007/demo/k8s/internal"

	"github.com/gin-gonic/gin"

	_ "github.com/liuqianhong6007/demo/k8s/api"
	"github.com/liuqianhong6007/demo/k8s/com"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "./config.yaml", "config file path")
	flag.Parse()
}

func main() {
	internal.ReadConfig(cfgPath)

	// init k8s client
	internal.InitK8sClientset(internal.Cfg().Inner, internal.Cfg().CfgPath)

	r := gin.Default()
	com.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", internal.Cfg().Host, internal.Cfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
