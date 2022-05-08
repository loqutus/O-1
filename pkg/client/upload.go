package client

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// Upload uploads a file to the server.
func Upload(fileName string, path string, justWrite bool) error {
	fileNameWithoutPath := filepath.Base(fileName)
	logrus.Println("Uploading file: ", fileName)
	logrus.Println("Path: ", path)
	logrus.Println("Server:", types.Client.HostName+":"+types.Client.Port)
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	hostName := types.Client.HostName
	port := types.Client.Port
	url := ""
	if path[len(path)-1] == '/' {
		url = "http://" + hostName + ":" + port + "/" + path + fileNameWithoutPath
	} else {
		url = "http://" + hostName + ":" + port + "/" + path
	}
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
