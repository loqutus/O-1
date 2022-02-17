package client

import (
	"fmt"
	"net/http"
	"time"

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
}
