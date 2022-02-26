package main

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: o1 <command> <filename> <path> [<args>]")
	}
	cmd := os.Args[1]
	var HostName string
	flag.StringVar(&HostName, "host", "localhost", "hostname")
	var Port string
	flag.StringVar(&Port, "port", "6969", "port")
	var Timeout string
	flag.StringVar(&Timeout, "timeout", "10", "timeout in seconds")
	flag.Parse()

	types.Client.HostName = HostName
	types.Client.Port = Port
	timeoutInt, err := strconv.Atoi(Timeout)
	if err != nil {
		timeoutInt = 10
	}
	types.Client.Timeout = time.Second * time.Duration(timeoutInt)
	switch cmd {
	case "upload":
		if len(os.Args) < 4 {
			panic("Usage: o1 upload <filename> <path> [<args>]")
		}
		err := client.Upload(os.Args[2], os.Args[3], false)
		if err != nil {
			logrus.Fatal(err)
		}
	case "download":
		if len(os.Args) < 3 {
			panic("Usage: o1 download <filename> [<args>]")
		}
		err := client.Download(os.Args[2])
		if err != nil {
			logrus.Fatal(err)
		}
	case "delete":
		if len(os.Args) < 3 {
			panic("Usage: o1 delete <path> [<args>]")
		}
		err := client.Delete(os.Args[2], false)
		if err != nil {
			logrus.Fatal(err)
		}
	case "info":
		err := client.Info()
		if err != nil {
			logrus.Fatal(err)
		}
	default:
		panic("unknown command")
	}
	return
}
