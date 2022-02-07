package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func PostFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	fileNameWithPath := types.Server.LocalDir + "/" + fileName
	logrus.Println("PostFile " + fileName)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(err, w)
		return
	}
	defer r.Body.Close()
	logrus.Println("Writing file " + fileName)
	err = os.WriteFile(fileNameWithPath, body, 0644)
	if err != nil {
		Error(err, w)
		return
	}
	fileSize, fileHash, err := file.Write(fileNameWithPath, body)
	if err != nil {
		Error(err, w)
		return
	}
	nodes := chooseNodes()
	fileInfo := types.FileInfo{
		Name:   fileName,
		Size:   fileSize,
		SHA256: string(fileHash),
		Nodes:  nodes,
	}
	logrus.Println("Putting fileInfo to ETCD: " + fileName)
	fileInfoJSON, err := json.Marshal(fileInfo)
	if err != nil {
		Error(err, w)
		return
	}
	err = etcdclient.Put(fileName, string(fileInfoJSON))
	if err != nil {
		Error(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
