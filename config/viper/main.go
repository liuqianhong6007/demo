package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/config/viper/config"
)

func main() {
	os.Setenv("GAME_SERVICE_SERVER_ID", "20")
	os.Setenv("GAME_SERVICE_PORT", "8080")

	config.ReadConf()

	fmt.Println("ServiceID: ", config.ServiceId())
	fmt.Println("ServicePort: ", config.ServicePort())
	fmt.Println("PassportKey: ", config.PassportKey())
	fmt.Println("GameplayManagerAddr: ", config.GameplayManagerAddr())
	fmt.Println("EtcdEndpoint: ", config.EtcdEndpoint())

	fmt.Println("FastwayAddr: ", config.FastwayAddr())
	fmt.Println("FastwayAuthkey: ", config.FastwayAuthkey())

	fmt.Println("DBHost: ", config.DBHost())
	fmt.Println("DBPort: ", config.DBPort())
	fmt.Println("DBUser: ", config.DBUser())
	fmt.Println("DBPassword: ", config.DBPassword())
	fmt.Println("DBName: ", config.DBName())
	fmt.Println("LogDir: ", config.LogDir())
	fmt.Println("LogName: ", config.LogName())
	fmt.Println("LogMaxSize: ", config.LogMaxSize())
	fmt.Println("LogMaxAge: ", config.LogMaxAge())
	fmt.Println("LogMaxBackups: ", config.LogMaxBackups())
	fmt.Println("LogCompress: ", config.LogCompress())
	fmt.Println("LogChanSize: ", config.LogChanSize())
	fmt.Println("LogReleaseMode: ", config.LogReleaseMode())

	config.ReadFromEtcd(config.EtcdEndpoint(), "/conf/game_service")
}
