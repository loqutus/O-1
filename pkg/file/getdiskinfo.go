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
	types.DiskInfo.DiskFree = humanize.Bytes(di.Free)
	types.DiskInfo.DiskUsed = humanize.Bytes(di.Used)
	types.DiskInfo.DiskTotal = humanize.Bytes(di.Total)
	return nil
}
