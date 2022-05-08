package restapi

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

// DeleteFile handles file removal requests.
func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	fileNameWithPath := types.Server.LocalDir + "/" + fileName
	logrus.Println("DeleteFile " + fileNameWithPath)
	err := os.Remove(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	fileInfoString, err := etcdclient.Get(fileName)
	if err != nil {
		Error(err, w)
		return
	}
	types.Client.Port = types.Server.ListenPort
	var fileInfo types.FileInfo
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		Error(err, w)
		return
	}
	justDelete := false
	justDeleteString := r.Header.Get("O1-Just-Delete")
	if justDeleteString == "true" {
		justDelete = true
	}
	if !justDelete {
		for _, node := range fileInfo.Nodes {
			if node == types.Server.HostName {
				continue
			}
			types.Client.HostName = node
			err := client.Delete(fileName, true)
			if err != nil {
				Error(err, w)
				return
			}
		}
		err = etcdclient.Delete(fileName)
		if err != nil {
			Error(err, w)
			return
		}
	}
	types.Info.FilesCount--
	types.Info.Used -= uint64(fileInfo.Size)
	types.Info.Free += uint64(fileInfo.Size)
	w.WriteHeader(http.StatusOK)
}
