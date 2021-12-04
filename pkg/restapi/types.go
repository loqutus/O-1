import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type RestAPI struct {
	LocalDir string
	Ctx      *context.Context
	Cli      clientv3.Client
}