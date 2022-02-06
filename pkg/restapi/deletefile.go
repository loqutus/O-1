package restapi

import (
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	fileNameWithPath := types.Server.LocalDir + "/" + fileName
	logrus.Println("DeleteFile " + fileNameWithPath)
	err := os.Remove(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	err = etcdclient.Delete(fileName)
	if err != nil {
		Error(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
