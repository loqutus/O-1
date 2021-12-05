package restapi

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/types"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path
	fileInfoString, err := api.Cli.Get(api.Ctx, fileName)
	if err != nil {
		Error(err, &w)
		return
	}
	var fileInfo types.FileInfo
	err := json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		Error(err, &w)
		return
	}
	fileBody, err := os.ReadFile(api.LocalDir + "/" + fileName)
	if err != nil {
		Error(err, &w)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBody)
}
