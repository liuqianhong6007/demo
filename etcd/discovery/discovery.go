package discovery

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type Discovery struct {
	client          *clientv3.Client
	leaseGrants     map[string]*clientv3.LeaseGrantResponse
	leaseGrantMutex sync.Mutex
	logger          *zap.Logger
	closeChan       chan struct{}
}

func NewDiscovery(endpoints string, logger *zap.Logger) *Discovery {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return &Discovery{
		client:      client,
		leaseGrants: make(map[string]*clientv3.LeaseGrantResponse),
		logger:      logger,
		closeChan:   make(chan struct{}),
	}
}

func (s *Discovery) Register(key, val string) error {
	s.leaseGrantMutex.Lock()
	defer s.leaseGrantMutex.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	leaseGrant, err := s.client.Grant(ctx, 5)
	if err != nil {
		return err
	}

	_, err = s.client.Put(ctx, key, val, clientv3.WithLease(leaseGrant.ID))
	if err != nil {
		return err
	}

	/*keepAliveCtx, keepAliveCancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer keepAliveCancel()*/

	leaseKeepActive, err := s.client.KeepAlive(context.Background(), leaseGrant.ID)
	if err != nil {
		return err
	}

	s.leaseGrants[key] = leaseGrant
	s.logger.Info(fmt.Sprintf("register success: key = %s", key))

	s.doAsync(func() {
		for {
			select {
			case _, ok := <-leaseKeepActive:
				if !ok {
					s.logger.Info("etcd keep alive channel closed")
					for {
						select {
						case <-time.After(time.Second):
							s.UnRegister(key)
							if err = s.Register(key, val); err == nil {
								return
							}
						case <-s.closeChan:
							s.logger.Info("exit watch loop goroutine")
							return
						}
					}
				}
			}
		}
	})

	return nil
}

func (s *Discovery) UnRegister(key string) error {
	s.leaseGrantMutex.Lock()
	defer s.leaseGrantMutex.Unlock()

	if leaseGrant, ok := s.leaseGrants[key]; ok {
		delete(s.leaseGrants, key)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := s.client.Revoke(ctx, leaseGrant.ID)
		if err != nil {
			s.logger.Error("revoke error", zap.String("key", key), zap.Error(err))
			return err
		}
	}
	return nil
}

func (s *Discovery) Watch(keyPrefix string) {
	// 开启监控
	s.doAsync(func() {
		watchChan := s.client.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
		for {
			select {
			case event, ok := <-watchChan:
				if !ok || event.Canceled {
					s.logger.Error("[watch] channel closed or canceled", zap.String("key", keyPrefix))
					return
				}
				if event.Created {
					continue
				}

				for _, ev := range event.Events {
					s.logger.Info("[watch] receive event", zap.Int("type", int(ev.Type)), zap.String("key", string(ev.Kv.Key)), zap.String("val", string(ev.Kv.Value)))
				}
			}
		}
	})

	// 获取初始状态
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsp, err := s.client.Get(ctx, keyPrefix, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	for _, kv := range rsp.Kvs {
		s.logger.Info("[watch] get", zap.String("key", string(kv.Key)), zap.String("val", string(kv.Value)))
	}
}

func (s *Discovery) Stop() {
	s.leaseGrantMutex.Lock()
	defer s.leaseGrantMutex.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, leaseGrant := range s.leaseGrants {
		_, err := s.client.Revoke(ctx, leaseGrant.ID)
		if err != nil {
			s.logger.Error("revoke error", zap.Int64("leaseID", int64(leaseGrant.ID)))
		}
	}
	close(s.closeChan)
	s.client.Close()
}

func (s *Discovery) doAsync(work func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s.logger.Error("panic happens", zap.Any("err", err))
			}
		}()
		work()
	}()
}

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}
