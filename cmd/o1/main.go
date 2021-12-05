package main

import (
	"os"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	localDir := os.Getenv("O1_LOCAL_DIR")
	if localDir == "" {
		localDir = "/tmp/O-1"
	}
	ctx, cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	server.Start(&ctx, cli, localDir)
}
