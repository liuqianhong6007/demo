package main

import (
	"context"

	"github.com/liuqianhong6007/demo/auth/api"
	_ "github.com/liuqianhong6007/demo/auth/api"
	"github.com/liuqianhong6007/demo/auth/config"
	"github.com/liuqianhong6007/demo/auth/internal"
)

func main() {
	// 读取配置
	config.ReadConf()

	// 初始化数据库
	internal.InitDatabase(config.DbConf())

	// 检查数据库
	ctx := context.Background()
	api.CheckContext(ctx)

	// 开启 http 服务
	internal.Start(config.Host(), config.Port())
}
