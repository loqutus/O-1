package etcdclient

import (
	"context"
	"time"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func New() (context.Context, *clientv3.Client, error) {
	logrus.Println("etcd client init")
	ctx := context.Background()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + types.Server.ETCDHost + ":" + types.Server.ETCDPort},
		DialTimeout: 10 * time.Second,
	})
	return ctx, cli, err
}
