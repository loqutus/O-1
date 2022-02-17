package restapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
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
	var f *os.File
	fileBody := []byte{}
	if fileHere {
		logrus.Println("File", fileName, "should be here")
		f, err = os.Open(filepath.Join(types.Server.LocalDir, fileName))
		//fileBody, err = os.ReadFile(filepath.Join(types.Server.LocalDir, fileName))
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
				f, err = os.Open(filepath.Join(types.Server.LocalDir, fileName))
				if err != nil {
					Error(err, w)
					return
				}
				err = os.Remove(filepath.Join(types.Server.LocalDir, fileName))
				if err != nil {
					Error(err, w)
				}
			} else {
				Error(err, w)
				return
			}
		}
	} else {
		logrus.Println("File", fileName, "should be elsewhere")
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

	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
	w.Write(fileBody)
}
