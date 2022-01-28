package etcdclient

import (
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func New() (*clientv3.Client, error) {
	logrus.Println("etcd client init")
	logrus.Println("ETCDHost: ", types.Server.ETCDHost)
	logrus.Println("ETCDPort: ", types.Server.ETCDPort)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + types.Server.ETCDHost + ":" + types.Server.ETCDPort},
		DialTimeout: types.Server.Timeout,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	return cli, err
}
