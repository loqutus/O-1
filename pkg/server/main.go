package server

import (
	"context"
	"github.com/loqutus/O-1/pkg/restapi"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Start(ctx *context.Context, cli *clientv3.Client) {
	restapi.Start(ctx, cli)
}
