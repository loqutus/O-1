package restapi

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	logrus.Println("GetFile " + fileName)
	fileInfo, err := GetFileInfo(fileName)
	if err != nil {
		Error(err, w)
		return
	}
	fileShouldBeHere := CheckIfFileShouldBeHere(types.Server.HostName, fileInfo.Nodes)
	var f *os.File
	defer f.Close()
	if fileShouldBeHere {
		logrus.Println("File", fileName, "should be here")
		f, err = os.Open(filepath.Join(types.Server.LocalDir, fileName))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				logrus.Println("File is not found, trying to get it from another nodes")
				err = getFileFromNodes(fileName, fileInfo.Nodes)
				if err != nil {
					Error(err, w)
					return
				}
				err = MoveFile(fileName)
				if err != nil {
					Error(err, w)
					return
				}
				f, err = os.Open(filepath.Join(types.Server.LocalDir, fileName))
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
		err = getFileFromNodes(fileName, fileInfo.Nodes)
		if err != nil {
			Error(err, w)
			return
		}
		file.EnsureDir(filepath.Join(types.Server.LocalDir, filepath.Dir(fileName)))
		err = MoveFile(fileName)
		if err != nil {
			Error(err, w)
			return
		}
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
}
