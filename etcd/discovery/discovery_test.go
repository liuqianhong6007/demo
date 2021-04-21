package discovery

import (
	"testing"

	"go.uber.org/goleak"
)

var (
	endpoints = "localhost:12379"
)

func TestDiscovery_Watch(t *testing.T) {
	d := NewDiscovery(endpoints, NewZapLogger())

	// 服务发现
	d.Watch("service")

	// 服务注册
	if err := d.Register("service/test", "hahahaha"); err != nil {
		panic(err)
	}

	d.Stop()
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
