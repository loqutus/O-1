package etcdclient

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func New() (context.Context, clientv3.Client, error) {
	ctx := context.Background()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://etcd:2379"},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return ctx, clientv3., err
	}
	return ctx, cli, nil
}
