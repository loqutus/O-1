package client

import (
	"errors"
	"net/http"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Delete(fileName string, justDelete bool) error {
	logrus.Println("Deleting file: ", fileName)
	hostname := types.Client.HostName
	port := types.Client.Port
	url := "http://" + hostname + ":" + port + "/" + fileName
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	if justDelete {
		req.Header.Set("O1-Just-Delete", "true")
	}
	client := http.Client{
		Timeout: types.Client.Timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
