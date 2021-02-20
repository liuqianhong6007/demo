package etcd

import (
	"context"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var (
	client  *clientv3.Client
	timeout = 5 * time.Second
)

func Init(endpoints string) {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic("Fatal error while new etcd client: " + err.Error())
	}
}

type Kv struct {
	Key string
	Val string
}

func Get(key string) ([]Kv, error) {
	ctx, cancleFunc := context.WithTimeout(context.TODO(), timeout)
	defer cancleFunc()

	rsp, err := client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var result []Kv
	for _, kv := range rsp.Kvs {
		result = append(result, Kv{
			Key: string(kv.Key),
			Val: string(kv.Value),
		})
	}
	return result, nil
}

func Put(key, value string) error {
	ctx, cancleFunc := context.WithTimeout(context.TODO(), timeout)
	defer cancleFunc()

	_, err := client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func Delete(key string) error {
	ctx, cancleFunc := context.WithTimeout(context.TODO(), timeout)
	defer cancleFunc()

	_, err := client.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
