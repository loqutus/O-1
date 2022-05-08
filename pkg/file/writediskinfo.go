package file

import (
	"encoding/json"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// WriteDiskInfo writes disk info to ETCD.
func WriteDiskInfo() error {
	logrus.Println("Updating disk info...")
	str, err := json.Marshal(types.Info)
	if err != nil {
		return err
	}
	nodeName := types.Server.NodeName
	err = etcdclient.Put(nodeName, string(str))
	return err
}
