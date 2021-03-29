package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/k8s/config"
	"github.com/liuqianhong6007/demo/k8s/internal"
)

func main() {
	config.ReadConfig("./config.yaml")

	// init k8s client
	internal.InitK8s(config.Cfg().K8s.CfgPath)

	r := gin.Default()
	internal.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", config.Cfg().Server.Host, config.Cfg().Server.Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
