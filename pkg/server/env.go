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
	logrus.Println("O1_LOCAL_DIR: ", localDir)
	nodeName := os.Getenv("O1_NODE_NAME")
	if nodeName == "" {
		nodeName = "localhost"
	}
	logrus.Println("O1_NODE_NAME: ", nodeName)
	listenPort := os.Getenv("O1_LISTEN_PORT")
	if listenPort == "" {
		listenPort = "6969"
	}
	logrus.Println("O1_LISTEN_PORT: ", listenPort)
	listenPortProbe := os.Getenv("O1_LISTEN_PORT_PROBE")
	if listenPortProbe == "" {
		listenPortProbe = "6970"
	}
	logrus.Println("O1_LISTEN_PORT_PROBE: ", listenPortProbe)
	etcdHost := os.Getenv("O1_ETCD_HOST")
	if etcdHost == "" {
		etcdHost = "localhost"
	}
	logrus.Println("env O1_ETCD_HOST: ", etcdHost)
	etcdPort := os.Getenv("O1_ETCD_PORT")
	if etcdPort == "" {
		etcdPort = "2379"
	}
	logrus.Println("env O1_ETCD_PORT: ", etcdPort)
	etcdUser := os.Getenv("O1_ETCD_USER")
	if etcdUser == "" {
		etcdUser = "root"
	}
	logrus.Println("env O1_ETCD_USER: ", etcdUser)
	etcdPassword := os.Getenv("O1_ETCD_PASSWORD")
	if etcdPassword == "" {
		etcdPassword = "root"
	}
	logrus.Println("env O1_ETCD_PASSWORD: ", etcdPassword)
	replicaCount := os.Getenv("O1_REPLICA_COUNT")
	logrus.Println("O1_REPLICA_COUNT: ", replicaCount)
	if replicaCount == "" {
		replicaCount = "1"
	}
	replicaCountInt, err := strconv.Atoi(replicaCount)
	if err != nil {
		logrus.Fatal(err)
	}
	serviceName := os.Getenv("O1_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "O1"
	}
	logrus.Println("O1_SERVICE_NAME: ", serviceName)
	namespace := os.Getenv("O1_NAMESPACE")
	if namespace == "" {
		namespace = "default"
	}
	logrus.Println("env O1_NAMESPACE: ", namespace)

	types.Server.LocalDir = localDir
	types.Server.NodeName = nodeName
	types.Server.Nodes = []string{nodeName}
	types.Server.ListenPort = listenPort
	types.Server.ListenPortProbe = listenPortProbe
	types.Server.ETCDHost = etcdHost
	types.Server.ETCDPort = etcdPort
	types.Server.ETCDUser = etcdUser
	types.Server.ETCDPassword = etcdPassword
	types.Server.ReplicaCount = replicaCountInt
	types.Server.ServiceName = serviceName
	types.Server.Timeout = 5 * time.Second
	types.Server.Namespace = namespace
	return
}
