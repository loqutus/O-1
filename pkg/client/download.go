package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(fileName string) error {
	resp, err := http.Get(fileName)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
