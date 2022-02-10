package client

import (
	"errors"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Upload(fileName string, justWrite bool) error {
	logrus.Println("Uploading file: ", fileName, "To", types.Client.HostName+":"+types.Client.Port)
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
	req, err := http.NewRequest("POST", url, f)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	if justWrite {
		req.Header.Set("O1-Just-Write", "true")
	} else {
		req.Header.Set("O1-Just-Write", "false")
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
