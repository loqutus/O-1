package etcdclient

import (
	"context"
	"errors"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// Get retrieves a key from ETCD.
func Get(key string) (string, error) {
	logrus.Println("ETCD: Get", key)
	ctx, cancel := context.WithTimeout(context.Background(), types.Server.Timeout)
	resp, err := types.Server.Cli.Get(ctx, key)
	cancel()
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", errors.New("file not found in ETCD")
	}
	return string(resp.Kvs[0].Value), nil
}
