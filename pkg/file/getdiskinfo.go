package file

import (
	humanize "github.com/dustin/go-humanize"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/minio/minio/internal/disk"
)

func GetDiskInfo() error {
	di, err := disk.GetInfo(types.Server.LocalDir)
	if err != nil {
		return err
	}
	types.Node.DiskFree = humanize.Bytes(di.Free)
	types.Node.DiskUsed = humanize.Bytes(di.Used)
	types.Node.DiskTotal = humanize.Bytes(di.Total)
	return nil
}
