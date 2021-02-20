package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/viper/config"
)

func main() {
	os.Setenv("GAME_SERVICE.SERVER_ID", "20")

	config.SetDefault(kServiceId, 1)
	config.SetDefault(kServicePort, 8000)
	config.SetDefault(kEtcdEndpoint, "http://127.0.0.1:2379")

	config.SetOverride(kLogDir, "log")

	config.ReadConf("game_service", "yaml")

	fmt.Printf("ServiceID: %s\n", ServiceId())
	fmt.Printf("ServicePort: %d\n", ServicePort())
	fmt.Printf("PassportKey: %s\n", PassportKey())
	fmt.Printf("GameplayManagerAddr: %s\n", GameplayManagerAddr())
	fmt.Printf("EtcdEndpoint: %s\n", EtcdEndpoint())

	fmt.Printf("FastwayAddr: %s\n", FastwayAddr())
	fmt.Printf("FastwayAuthkey: %s\n", FastwayAuthkey())

	fmt.Printf("DBHost: %s\n", DBHost())
	fmt.Printf("DBPort: %d\n", DBPort())
	fmt.Printf("DBUser: %s\n", DBUser())
	fmt.Printf("DBPassword: %s\n", DBPassword())
	fmt.Printf("DBName: %s\n", DBName())
	fmt.Printf("LogDir: %s\n", LogDir())
	fmt.Printf("LogName: %s\n", LogName())
	fmt.Printf("LogMaxSize: %d\n", LogMaxSize())
	fmt.Printf("LogMaxAge: %d\n", LogMaxAge())
	fmt.Printf("LogMaxBackups: %d\n", LogMaxBackups())
	fmt.Printf("LogCompress: %v\n", LogCompress())
	fmt.Printf("LogChanSize: %d\n", LogChanSize())
	fmt.Printf("LogReleaseMode: %v\n", LogReleaseMode())

	config.ReadFromEtcd(EtcdEndpoint(), "/conf/game_service")
}
