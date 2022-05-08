package types

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// FileInfo is the file info.
type FileInfo struct {
	Name   string
	Size   int64
	SHA256 string
	Nodes  []string
}

// ServerInfo is the server info.
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

// DiskInfo is the disk info.
type DiskInfo struct {
	FilesCount int
	Used       uint64
	Free       uint64
	Total      uint64
}

var Info DiskInfo

var Server ServerInfo

// ClientInfo is the client info.
type ClientInfo struct {
	HostName string
	Port     string
	Timeout  time.Duration
}

var Client ClientInfo
