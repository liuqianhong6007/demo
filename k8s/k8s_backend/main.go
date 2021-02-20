package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/k8s_backend/config"
	"github.com/liuqianhong6007/k8s_backend/k8s"
	"github.com/liuqianhong6007/k8s_backend/route"
)

func main() {
	// init k8s client
	k8s.InitK8s(config.K8sCfg().CfgPath)

	r := gin.Default()
	route.RegisterRoute(r)

	serverAddr := fmt.Sprintf("%s:%d", config.ServerCfg().Host, config.ServerCfg().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
