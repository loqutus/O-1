package restapi

import (
	"math/rand"
	"time"

	"github.com/loqutus/O-1/pkg/types"
)

// cnooseNodes returns replicaCount random nodes from the nodes list
func chooseNodes() []string {
	if len(types.Server.Nodes) <= types.Server.ReplicaCount {
		return types.Server.Nodes
	} else {
		rand.Seed(time.Now().Unix())
		res := []string{}
		for i := 0; i < types.Server.ReplicaCount; i++ {
			n := rand.Intn(len(types.Server.Nodes))
			res = append(res, types.Server.Nodes[n])
		}
		return res
	}
}
