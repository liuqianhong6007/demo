package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/etcd/etcd_backend/conf_key"
	"github.com/liuqianhong6007/demo/etcd/etcd_backend/etcd"
	"github.com/liuqianhong6007/demo/etcd/etcd_backend/route"
)

func main() {
	// 初始化 etcd
	etcd.Init(conf_key.EtcdAddr())

	// 开启 http 服务
	r := gin.Default()
	route.RegisterRoute(r)
	serverAddr := fmt.Sprintf("%s:%d", conf_key.Host(), conf_key.Port())
	if err := r.Run(serverAddr); err != nil {
		panic(err)
	}
}
