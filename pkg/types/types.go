package types

type FileInfo struct {
	Name   string
	Size   int64
	SHA256 string
	Nodes  []string
}
