package server

import (
	"github.com/loqutus/O-1/pkg/restapi"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Start() {
	logrus.Println("Starting server...")
	EnsureDir(types.Server.LocalDir)
	restapi.Start()
}
