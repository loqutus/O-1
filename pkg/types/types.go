package types

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type FileInfo struct {
	Name   string
	Size   int64
	SHA256 string
	Nodes  []string
}

type ServerInfo struct {
	LocalDir   string
	NodeName   string
	Nodes      []string
	Cli        *clientv3.Client
	ListenPort string
	ETCDHost   string
	ETCDPort   string
	Timeout    time.Duration
}

var Server ServerInfo

type ClientInfo struct {
	HostName string
	Port     string
}

var Client ClientInfo
