package main

import (
	"os"
	"time"

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
	localDir := os.Getenv("O1_LOCAL_DIR")
	if localDir == "" {
		localDir = "/tmp/O1"
	}
	logrus.Println("env O1_LOCAL_DIR: ", localDir)
	nodeName := os.Getenv("O1_NODE_NAME")
	if nodeName == "" {
		nodeName = "localhost"
	}
	listenPort := os.Getenv("O1_LISTEN_PORT")
	logrus.Println("env O1_LISTEN_PORT: ", listenPort)
	if listenPort == "" {
		listenPort = "6969"
	}
	etcdHost := os.Getenv("O1_ETCD_HOST")
	logrus.Println("env O1_ETCD_HOST: ", etcdHost)
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	etcdPort := os.Getenv("O1_ETCD_PORT")
	logrus.Println("env O1_ETCD_PORT: ", etcdPort)
	if etcdPort == "" {
		etcdPort = "2379"
	}
	etcdUser := os.Getenv("O1_ETCD_USER")
	logrus.Println("env O1_ETCD_USER: ", etcdUser)
	if etcdUser == "" {
		etcdUser = "root"
	}
	etcdPassword := os.Getenv("O1_ETCD_PASSWORD")
	logrus.Println("env O1_ETCD_PASSWORD: ", etcdPassword)
	if etcdPassword == "" {
		etcdPassword = "root"
	}

	types.Server.LocalDir = localDir
	types.Server.NodeName = nodeName
	types.Server.Nodes = []string{nodeName}
	types.Server.ListenPort = listenPort
	types.Server.ETCDHost = etcdHost
	types.Server.ETCDPort = etcdPort
	types.Server.ETCDUser = etcdUser
	types.Server.ETCDPassword = etcdPassword
	types.Server.Timeout = 5 * time.Second

	cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	types.Server.Cli = cli
	defer types.Server.Cli.Close()
	server.Start()
}
