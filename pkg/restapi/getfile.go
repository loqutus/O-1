package restapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	logrus.Println("GetFile " + fileName)
	fileInfoString, err := etcdclient.Get(fileName)
	if err != nil {
		Error(err, w)
		return
	}
	fileHere := false
	var fileInfo types.FileInfo
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		Error(err, w)
		return
	}
	for _, nodeName := range fileInfo.Nodes {
		if nodeName == types.Server.HostName {
			fileHere = true
			break
		}
	}
	fileBody := []byte{}
	if fileHere {
		logrus.Println("File", fileName, "should be here")
		fileBody, err = os.ReadFile(filepath.Join(types.Server.LocalDir, fileName))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				err = getFileFromNodes(fileName, fileInfo.Nodes)
				if err != nil {
					Error(err, w)
					return
				}
				file.EnsureDir(filepath.Join(types.Server.LocalDir, filepath.Dir(fileName)))
				err = os.Rename(fileName, filepath.Join(types.Server.LocalDir, fileName))
				if err != nil {
					Error(err, w)
					return
				}
				fileBody, err = os.ReadFile(filepath.Join(types.Server.LocalDir, fileName))
				if err != nil {
					Error(err, w)
					return
				}
			} else {
				Error(err, w)
				return
			}
		}
	} else {
		logrus.Println("File", fileName, "should be elsewhere")

	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBody)
}
