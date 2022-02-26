package file

import (
	"time"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func InfoWatcher() {
	previousFileCount := types.Info.FilesCount
	for {
		if previousFileCount != types.Info.FilesCount {
			previousFileCount = types.Info.FilesCount
			err := WriteDiskInfo()
			if err != nil {
				logrus.Println(err)
			}
		}
		time.Sleep(types.Server.Timeout)
	}
}
