package main

import (
	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/server"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	logrus.SetReportCaller(true)
	logrus.Println("O-1 server starting...")
	server.ParseEnv()
	cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	types.Server.Ready = true
	types.Server.Cli = cli
	defer types.Server.Cli.Close()
	server.Start()
}
