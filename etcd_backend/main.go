package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/etcd_backend/config"
	"github.com/liuqianhong6007/demo/etcd_backend/etcd"
	"github.com/liuqianhong6007/demo/etcd_backend/route"
)

func main() {
	// 读取配置
	config.ReadConf()

	// 初始化 etcd
	etcd.Init(config.GetConfig().EtcdAddr)

	// 开启 http 服务
	r := gin.Default()
	route.RegisterRoute(r)
	serverAddr := fmt.Sprintf("%s:%d", config.GetConfig().Host, config.GetConfig().Port)
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
