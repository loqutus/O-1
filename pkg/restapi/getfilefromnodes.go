package restapi

import (
	"errors"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// GetFileFromNodes gets a file from other nodes.
func getFileFromNodes(fileName string, nodes []string) error {
	downloaded := false
	for _, node := range nodes {
		logrus.Println("Trying node", node)
		types.Client.HostName = node
		types.Client.Port = types.Server.ListenPort
		err := client.Download(fileName)
		if err != nil {
			continue
		} else {
			downloaded = true
			logrus.Println("Found file", fileName, "at node", node)
			break
		}
	}
	if !downloaded {
		return errors.New("File " + fileName + " not found on other nodes")
	} else {
		return nil
	}
}
