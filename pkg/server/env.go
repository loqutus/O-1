package server

import (
	"os"
	"strconv"
	"time"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func ParseEnv() {
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
	replicaCount := os.Getenv("O1_REPLICA_COUNT")
	logrus.Println("env O1_REPLICA_COUNT: ", replicaCount)
	if replicaCount == "" {
		replicaCount = "1"
	}
	replicaCountInt, err := strconv.Atoi(replicaCount)
	if err != nil {
		logrus.Fatal(err)
	}
	serviceName := os.Getenv("O1_SERVICE_NAME")
	logrus.Println("env O1_SERVICE_NAME: ", serviceName)
	if serviceName == "" {
		serviceName = "O1"
	}
	namespace := os.Getenv("O1_NAMESPACE")
	logrus.Println("env O1_NAMESPACE: ", namespace)
	if namespace == "" {
		namespace = "default"
	}
	types.Server.LocalDir = localDir
	types.Server.NodeName = nodeName
	types.Server.Nodes = []string{nodeName}
	types.Server.ListenPort = listenPort
	types.Server.ETCDHost = etcdHost
	types.Server.ETCDPort = etcdPort
	types.Server.ETCDUser = etcdUser
	types.Server.ETCDPassword = etcdPassword
	types.Server.ReplicaCount = replicaCountInt
	types.Server.ServiceName = serviceName
	types.Server.Timeout = 5 * time.Second
	types.Server.Namespace = namespace
}
