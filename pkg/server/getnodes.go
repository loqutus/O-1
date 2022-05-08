package server

import (
	"os"
	"strconv"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// getNodes gets a list of nodes.
func getNodes() {
	res := []string{}
	serviceName := types.Server.ServiceName
	for i := 0; i < types.Server.ReplicaCount; i++ {
		podNum := strconv.Itoa(i)

		node := serviceName + "-" + podNum + "." + serviceName + "." + types.Server.Namespace + ".svc.cluster.local"
		res = append(res, node)
		logrus.Println("Node added: ", node)
	}
	types.Server.Nodes = res
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostName = hostName + "." + serviceName + "." + types.Server.Namespace + ".svc.cluster.local"
	logrus.Println("Hostname:", hostName)
	types.Server.HostName = hostName
}
