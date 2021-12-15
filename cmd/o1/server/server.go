package main

import (
	"os"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/server"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	logrus.Println("O-1 server starting...")
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	localDir := os.Getenv("O1_LOCAL_DIR")
	if localDir == "" {
		localDir = "/tmp/O1"
	}
	nodeName := os.Getenv("O1_NODE_NAME")
	if nodeName == "" {
		nodeName = "localhost"
	}
	listenPort := os.Getenv("O1_LISTEN_PORT")
	if listenPort == "" {
		listenPort = "6969"
	}
	etcdHost := os.Getenv("O1_ETCD_HOST")
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	etcdPort := os.Getenv("O1_ETCD_PORT")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	types.Server.LocalDir = localDir
	types.Server.NodeName = nodeName
	types.Server.Nodes = []string{nodeName}
	types.Server.ListenPort = listenPort
	types.Server.ETCDHost = etcdHost
	types.Server.ETCDPort = etcdPort
	ctx, cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	types.Server.Cli = cli
	types.Server.Ctx = &ctx
	server.Start()
}
