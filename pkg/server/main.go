package server

import (
	"time"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/restapi"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// Server is the main server.
func Start() {
	logrus.Println("Starting server...")
	go startProbe()
	getNodes()
	file.EnsureDir(types.Server.LocalDir)
	err := file.GetDiskInfo()
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		cli, err := etcdclient.New()
		if err != nil {
			logrus.Println(err)
			time.Sleep(types.Server.Timeout)
		} else {
			logrus.Println("Server is ready")
			types.Server.Ready = true
			types.Server.Cli = cli
			defer types.Server.Cli.Close()
			go file.InfoWatcher()
			restapi.Start()
			break
		}
	}
	return
}
