package restapi

import (
	"encoding/json"

	"github.com/loqutus/O-1/pkg/etcdclient"
	"github.com/loqutus/O-1/pkg/types"
)

func GetFileInfo(fileName string) (types.FileInfo, error) {
	fileInfoString, err := etcdclient.Get(fileName)
	if err != nil {
		return types.FileInfo{}, err
	}
	var fileInfo types.FileInfo
	err = json.Unmarshal([]byte(fileInfoString), &fileInfo)
	if err != nil {
		return types.FileInfo{}, err
	}
	return fileInfo, nil
}
