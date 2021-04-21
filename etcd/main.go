package main

import (
	"flag"
	"time"

	"github.com/liuqianhong6007/demo/etcd/discovery"
)

var (
	etcdEndpoints = flag.String("etcd_endpoints", "localhost:2379", "etcd endpoints")
)

func main() {
	d := discovery.NewDiscovery(*etcdEndpoints, discovery.NewZapLogger())

	// 服务发现
	d.Watch("service")

	// 服务注册
	if err := d.Register("service/test", "hahahaha"); err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Minute)
	}
}
