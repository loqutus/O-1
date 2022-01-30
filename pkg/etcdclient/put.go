package etcdclient

import (
	"context"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Put(key, data string) error {
	logrus.Println("ETCD: Put", key)
	ctx, cancel := context.WithTimeout(context.Background(), types.Server.Timeout)
	_, err := types.Server.Cli.Put(ctx, key, data)
	cancel()
	return err
}
