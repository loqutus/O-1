package client

import (
	"errors"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Upload(fileName string) error {
	logrus.Println("Uploading file: ", fileName)
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	hostname := types.Client.HostName
	port := types.Client.Port
	url := "http://" + hostname + ":" + port + "/" + fileName
	client := http.Client{
		Timeout: types.Client.Timeout,
	}
	resp, err := client.Post(url, "application/data", f)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
