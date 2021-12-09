package main

import (
	"flag"
	"os"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/types"
)

func main() {
	cmd := os.Args[1]
	var HostName string
	flag.StringVar(&HostName, "host", "localhost", "hostname")
	var Port string
	flag.StringVar(&Port, "port", "6969", "port")
	flag.Parse()
	types.Client.Hostname = HostName
	types.Client.Port = Port

	switch cmd {
	case "upload":
		client.Upload(os.Args[2])
	case "download":
		client.Download(os.Args[2])
	default:
		panic("unknown command")
	}
	return
}
