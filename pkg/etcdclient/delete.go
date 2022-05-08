package etcdclient

import (
	"context"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// Delete deletes a key from ETCD.
func Delete(key string) error {
	logrus.Println("ETCD: Delete", key)
	ctx, cancel := context.WithTimeout(context.Background(), types.Server.Timeout)
	_, err := types.Server.Cli.Delete(ctx, key)
	cancel()
	return err
}
