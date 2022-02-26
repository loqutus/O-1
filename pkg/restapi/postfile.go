package restapi

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/sha256"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func PostFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	fileNameWithPath := types.Server.LocalDir + "/" + fileName
	logrus.Println("PostFile " + fileName)
	defer r.Body.Close()
	file.EnsureDir(filepath.Join(types.Server.LocalDir + filepath.Dir(fileName)))
	file, err := os.Create(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	fileSize, err := io.Copy(file, r.Body)
	if err != nil {
		Error(err, w)
		return
	}
	fileHash, err := sha256.GetFileSHA256(fileNameWithPath)
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
	types.Info.FilesCount++
	types.Info.Used += uint64(fileSize)
	types.Info.Free -= uint64(fileSize)
	w.WriteHeader(http.StatusOK)
}
