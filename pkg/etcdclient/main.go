package etcdclient

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func New() (context.Context, *clientv3.Client, error) {
	logrus.Println("etcd client init")
	ctx := context.Background()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://etcd:2379"},
		DialTimeout: 10 * time.Second,
	})
	return ctx, cli, err
}
