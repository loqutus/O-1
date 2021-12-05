package restapi

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/types"
)

func PostFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(err, &w)
		return
	}
	err = os.WriteFile(api.Path+"/"+fileName, body, 0644)
	if err != nil {
		Error(err, &w)
		return
	}
	fi, err := os.Stat(api.Path + "/" + fileName)
	if err != nil {
		Error(err, &w)
		return
	}
	fileSize := fi.Size()
	hash := sha256.New()
	if _, err := io.Copy(hash, r.Body); err != nil {
		Error(err, &w)
		return
	}
	SHA256 := hash.Sum(nil)
	fileInfo := types.FileInfo{
		Name:   fileName,
		Size:   fileSize,
		SHA256: string(SHA256),
		Nodes:  []string{},
	}
	fileNameJSON, _ := json.Marshal(fileInfo)
	_, err = api.Cli.Put(api.Ctx, fileNameJSON, fileInfo)
	if err != nil {
		Error(err, &w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
