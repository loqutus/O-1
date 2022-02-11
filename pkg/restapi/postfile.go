package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/loqutus/O-1/pkg/client"
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
	file.EnsureDir(filepath.Join(types.Server.LocalDir + filepath.Dir(fileName)))
	fileSize, fileHash, err := file.Write(fileNameWithPath, body)
	if err != nil {
		Error(err, w)
		return
	}
	nodes := chooseNodes()
	types.Client.Port = types.Server.ListenPort
	types.Client.Timeout = 5 * time.Second
	justWrite := false
	justWriteString := r.Header.Get("O1-Just-Write")
	if justWriteString == "true" {
		justWrite = true
	}
	if !justWrite {
		for _, node := range nodes {
			if node == types.Server.HostName {
				continue
			}
			types.Client.HostName = node
			err := client.Upload(fileNameWithPath, fileName, true)
			if err != nil {
				Error(err, w)
				return
			}
		}
	}
	fileInfo := types.FileInfo{
		Name:   fileName,
		Size:   fileSize,
		SHA256: string(fileHash),
		Nodes:  nodes,
	}
	if !justWrite {
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
	}
	w.WriteHeader(http.StatusOK)
}
