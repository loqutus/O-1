package etcdclient

import (
	"crypto/tls"
	"os"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// New returns new ETCD client.
func New() (*clientv3.Client, error) {
	logrus.Println("ETCD client init")
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + types.Server.ETCDHost + ":" + types.Server.ETCDPort},
		Username:    types.Server.ETCDUser,
		Password:    types.Server.ETCDPassword,
		DialTimeout: types.Server.Timeout,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		TLS:         &tls.Config{InsecureSkipVerify: true},
	})
	return cli, err
}
