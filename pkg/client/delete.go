package client

import (
	"errors"
	"net/http"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Delete(fileName string) error {
	logrus.Println("Deleting file: ", fileName)
	hostname := types.Client.HostName
	port := types.Client.Port
	url := "http://" + hostname + ":" + port + "/" + fileName
	req, err := http.NewRequest("DELETE", url, nil)
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
