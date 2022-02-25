package file

import (
	"encoding/json"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/types"
)

func WriteDiskInfo() error {
	str, err := json.Marshal(types.DiskInfo)
	if err != nil {
		return err
	}
	nodeName := types.Server.NodeName
	err = etcdclient.Put(nodeName, string(str))
	return err
}
