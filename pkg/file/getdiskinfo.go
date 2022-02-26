package file

import (
	"github.com/loqutus/O-1/pkg/types"
	"github.com/shirou/gopsutil/disk"
)

func GetDiskInfo() error {
	usage, err := disk.Usage(types.Server.LocalDir)
	if err != nil {
		return err
	}
	types.DiskInfo.Used = usage.Used
	types.DiskInfo.Free = usage.Free
	types.DiskInfo.Total = usage.Total
	return nil
}
