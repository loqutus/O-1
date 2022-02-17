package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Download(fileName string) error {
	logrus.Println("Downloading file: ", fileName)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := "http://" + types.Client.HostName + ":" + types.Client.Port + "/" + fileName
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	f, err := os.Create(filepath.Base(fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
