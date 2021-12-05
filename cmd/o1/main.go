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
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	localDir := os.Getenv("O1_LOCAL_DIR")
	if localDir == "" {
		localDir = "/tmp/O-1"
	}
	nodeName := os.Getenv("O1_NODE_NAME")
	if nodeName == "" {
		nodeName = "localhost"
	}

	types.Server.LocalDir = localDir
	types.Server.NodeName = nodeName
	types.Server.Nodes = []string{nodeName}

	ctx, cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	types.ServerInfo.Cli = cli
	types.ServerInfo.Ctx = ctx
	server.Start()
}
