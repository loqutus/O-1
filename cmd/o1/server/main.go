package main

import (
	"github.com/loqutus/O-1/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	logrus.SetReportCaller(true)
	logrus.Println("O-1 server starting...")
	server.ParseEnv()
	server.Start()
}
