package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/loqutus/O-1/pkg/types"
)

func Info() error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := "http://" + types.Client.HostName + ":" + types.Client.Port + "/info"
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	allBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var diskInfo types.DiskInfo
	err = json.Unmarshal(allBody, &diskInfo)
	if err != nil {
		return err
	}
	fmt.Println("Disk Used:", humanize.Bytes(diskInfo.Used))
	fmt.Println("Disk Free:", humanize.Bytes(diskInfo.Free))
	fmt.Println("Disk Total:", humanize.Bytes(diskInfo.Total))
	fmt.Println("Files Count:", diskInfo.FilesCount)
	return nil
}
