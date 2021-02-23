package main

import (
	"fmt"
	"os"

	"github.com/liuqianhong6007/demo/config/koanf/config"
)

func main() {
	os.Setenv("GAME_SERVICE.SERVER_ID", "20")

	config.SetEnvPrefix("")
	config.ReadConf("demo", "yml")
	fmt.Println("ServiceID: ", config.Get("game_service.server_id"))
	fmt.Println("ServicePort: ", config.Get("game_service.port"))
	fmt.Println("PassportKey: ", config.Get("game_service.passport_key"))
	fmt.Println("GameplayManagerAddr: ", config.Get("game_service.gameplay_manager_addr"))
	fmt.Println("EtcdEndpoint: ", config.Get("game_service.etcd_endpoint"))
}
