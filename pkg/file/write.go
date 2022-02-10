package file

import (
	"os"
	"strings"

	"github.com/loqutus/O-1/pkg/sha256"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func Write(fileNameWithPath string, data []byte) (int64, string, error) {
	logrus.Println("Writing file:", fileNameWithPath)
	fileName := strings.TrimPrefix(fileNameWithPath, types.Server.LocalDir)
	fileNameSplit := strings.Split(fileName, "/")
	fileNameSplitLen := len(fileNameSplit)
	fileNameDir := strings.Join(fileNameSplit[:fileNameSplitLen-1], "/")
	fullPathDir := types.Server.LocalDir + "/" + fileNameDir
	logrus.Println("fullPathDir:", fullPathDir)
	err := EnsureDir(fullPathDir)
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
