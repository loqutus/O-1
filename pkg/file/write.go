package file

import (
	"os"
	"strings"

	"github.com/loqutus/O-1/pkg/sha256"
)

func Write(fileNameWithPath string, data []byte) (int64, string, error) {
	err := EnsureDir(strings.Split(fileNameWithPath, "/")[-1])
	if err != nil {
		return 0, "", err
	}
	err = os.WriteFile(fileNameWithPath, data, 0644)
	if err != nil {
		return 0, "", err
	}
	fi, err := os.Stat(fileNameWithPath)
	if err != nil {
		return 0, "", err
	}
	fileSize := fi.Size()
	fileHash, err := sha256.GetFileSHA256(fileNameWithPath)
	if err != nil {
		return 0, "", err
	}
	return fileSize, fileHash, nil
}
