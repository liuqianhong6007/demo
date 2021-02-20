package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/viper/config"
)

func main() {
	os.Setenv("GAME_SERVICE.SERVER_ID", "20")

	config.ReadConf()

	fmt.Printf("ServiceID: %s\n", config.ServiceId())
	fmt.Printf("ServicePort: %d\n", config.ServicePort())
	fmt.Printf("PassportKey: %s\n", config.PassportKey())
	fmt.Printf("GameplayManagerAddr: %s\n", config.GameplayManagerAddr())
	fmt.Printf("EtcdEndpoint: %s\n", config.EtcdEndpoint())

	fmt.Printf("FastwayAddr: %s\n", config.FastwayAddr())
	fmt.Printf("FastwayAuthkey: %s\n", config.FastwayAuthkey())

	fmt.Printf("DBHost: %s\n", config.DBHost())
	fmt.Printf("DBPort: %d\n", config.DBPort())
	fmt.Printf("DBUser: %s\n", config.DBUser())
	fmt.Printf("DBPassword: %s\n", config.DBPassword())
	fmt.Printf("DBName: %s\n", config.DBName())
	fmt.Printf("LogDir: %s\n", config.LogDir())
	fmt.Printf("LogName: %s\n", config.LogName())
	fmt.Printf("LogMaxSize: %d\n", config.LogMaxSize())
	fmt.Printf("LogMaxAge: %d\n", config.LogMaxAge())
	fmt.Printf("LogMaxBackups: %d\n", config.LogMaxBackups())
	fmt.Printf("LogCompress: %v\n", config.LogCompress())
	fmt.Printf("LogChanSize: %d\n", config.LogChanSize())
	fmt.Printf("LogReleaseMode: %v\n", config.LogReleaseMode())
}
