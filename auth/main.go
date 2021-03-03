package main

import (
	"context"

	"github.com/liuqianhong6007/demo/auth/api"
	_ "github.com/liuqianhong6007/demo/auth/api"
	"github.com/liuqianhong6007/demo/auth/internal"
)

func main() {
	// 读取配置
	internal.ReadConf()

	// 初始化数据库
	dbConf := internal.DatabaseConf{
		Driver:      internal.DbDriver(),
		Host:        internal.DbHost(),
		Port:        internal.DbPort(),
		User:        internal.DbUser(),
		Password:    internal.DbPassword(),
		Lib:         internal.DbLib(),
		MaxIdleConn: internal.DbMaxIdleConn(),
		MaxOpenConn: internal.DbMaxOpenConn(),
	}
	internal.InitDatabase(dbConf)

	// 检查数据库
	ctx := context.Background()
	api.CheckContext(ctx)

	// 初始化 castbin
	internal.InitCasbin(internal.CastbinConf{
		ModelPath:    internal.CastbinModelPath(),
		PolicyDriver: internal.CastbinPolicyDriver(),
		PolicyPath:   internal.CastbinPolicyPath(),
		DbConf:       dbConf,
	})

	// 开启 http 服务
	internal.Start(internal.Host(), internal.Port())
}
