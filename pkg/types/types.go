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
	LocalDir        string
	NodeName        string
	Nodes           []string
	Cli             *clientv3.Client
	ListenPort      string
	ListenPortProbe string
	ETCDHost        string
	ETCDPort        string
	ETCDUser        string
	ETCDPassword    string
	Timeout         time.Duration
	ReplicaCount    int
	ServiceName     string
	Namespace       string
	HostName        string
	Ready           bool
}

type NodeInfo struct {
	FilesNum  int
	DiskUsed  string
	DiskFree  string
	DiskTotal string
}

var Node NodeInfo

var Server ServerInfo

type ClientInfo struct {
	HostName string
	Port     string
	Timeout  time.Duration
}

var Client ClientInfo
