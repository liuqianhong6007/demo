package main

import (
	_ "github.com/liuqianhong6007/demo/auth/api"
	"github.com/liuqianhong6007/demo/auth/internal"
)

func main() {
	// 读取配置
	internal.ReadConf()

	// 初始化数据库
	dbCfg := internal.GetConfig().Db
	dbConf := internal.DatabaseConf{
		Driver:      dbCfg.Driver,
		Host:        dbCfg.Host,
		Port:        dbCfg.Port,
		User:        dbCfg.User,
		Password:    dbCfg.Password,
		Lib:         dbCfg.Lib,
		MaxIdleConn: dbCfg.MaxIdleConn,
		MaxOpenConn: dbCfg.MaxOpenConn,
	}
	internal.InitDatabase(dbConf)

	// 初始化 castbin
	castbinCfg := internal.GetConfig().Casbin
	internal.InitCasbin(internal.CastbinConf{
		ModelPath:    castbinCfg.ModelPath,
		PolicyDriver: castbinCfg.PolicyDriver,
		PolicyPath:   castbinCfg.PolicyPath,
		DbConf:       dbConf,
	})

	// 开启 http 服务
	serverCfg := internal.GetConfig().Server
	internal.Start(serverCfg.Host, serverCfg.Port)
}
