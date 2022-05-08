package file

import (
	"github.com/loqutus/O-1/pkg/types"
	"github.com/shirou/gopsutil/v3/disk"
)

// GetDiskInfo returns the disk information.
func GetDiskInfo() error {
	usage, err := disk.Usage(types.Server.LocalDir)
	if err != nil {
		return err
	}
	types.Info.Used = usage.Used
	types.Info.Free = usage.Free
	types.Info.Total = usage.Total
	return nil
}
