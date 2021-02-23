package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/config/viper/config"
)

func main() {
	os.Setenv("GAME_SERVICE.SERVER_ID", "20")

	config.SetDefault(kServiceId, 1)
	config.SetDefault(kServicePort, 8000)
	config.SetDefault(kEtcdEndpoint, "http://127.0.0.1:2379")

	config.SetOverride(kLogDir, "log")

	config.ReadConf("demo", "yaml")

	fmt.Println("ServiceID: ", ServiceId())
	fmt.Println("ServicePort: ", ServicePort())
	fmt.Println("PassportKey: ", PassportKey())
	fmt.Println("GameplayManagerAddr: ", GameplayManagerAddr())
	fmt.Println("EtcdEndpoint: ", EtcdEndpoint())

	fmt.Println("FastwayAddr: ", FastwayAddr())
	fmt.Println("FastwayAuthkey: ", FastwayAuthkey())

	fmt.Println("DBHost: ", DBHost())
	fmt.Println("DBPort: ", DBPort())
	fmt.Println("DBUser: ", DBUser())
	fmt.Println("DBPassword: ", DBPassword())
	fmt.Println("DBName: ", DBName())
	fmt.Println("LogDir: ", LogDir())
	fmt.Println("LogName: ", LogName())
	fmt.Println("LogMaxSize: ", LogMaxSize())
	fmt.Println("LogMaxAge: ", LogMaxAge())
	fmt.Println("LogMaxBackups: ", LogMaxBackups())
	fmt.Println("LogCompress: ", LogCompress())
	fmt.Println("LogChanSize: ", LogChanSize())
	fmt.Println("LogReleaseMode: ", LogReleaseMode())

	config.ReadFromEtcd(EtcdEndpoint(), "/conf/game_service")
}
