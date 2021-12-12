package main

import (
	"flag"
	"os"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: o1 <command> <filename> [<args>]")
	}
	cmd := os.Args[1]
	var HostName string
	flag.StringVar(&HostName, "host", "localhost", "hostname")
	var Port string
	flag.StringVar(&Port, "port", "6969", "port")
	flag.Parse()
	types.Client.HostName = HostName
	types.Client.Port = Port

	switch cmd {
	case "upload":
		err := client.Upload(os.Args[2])
		if err != nil {
			logrus.Fatal(err)
		}
	case "download":
		err := client.Download(os.Args[2])
		if err != nil {
			logrus.Fatal(err)
		}
	default:
		panic("unknown command")
	}
	return
}
