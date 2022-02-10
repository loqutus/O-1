package server

import (
	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/restapi"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Start() {
	logrus.Println("Starting server...")
	go startProbe()
	getNodes()
	file.EnsureDir(types.Server.LocalDir)
	restapi.Start()
	cli, err := etcdclient.New()
	if err != nil {
		logrus.Fatal(err)
	}
	types.Server.Ready = true
	types.Server.Cli = cli
	defer types.Server.Cli.Close()
}
