package client

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/loqutus/O-1/pkg/sha256"
	fileSHA256 "github.com/loqutus/O-1/pkg/sha256"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func TestUpload(t *testing.T) {
	for i := 0; i < 10; i++ {
		upload(t)
	}
}

func upload(t *testing.T) {
	types.Client.HostName = "localhost"
	types.Client.Port = "6969"
	file, err := os.CreateTemp("", "o-1-test*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	data := []byte(gofakeit.HackerPhrase())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.Write(data); err != nil {
		t.Fatal(err)
	}
	fileName := file.Name()
	if err := Upload(fileName, filepath.Base(fileName), false); err != nil {
		t.Fatal(err)
	}
	correctHash, err := fileSHA256.GetFileSHA256(fileName)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
	os.Remove(file.Name())
	err = Download(filepath.Base(fileName))
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fileName)
	downloadedHash, err := fileSHA256.GetFileSHA256(filepath.Base(fileName))
	if err != nil {
		t.Fatal(err)
	}
	if correctHash != downloadedHash {
		t.Fatalf("hash mismatch: %s != %s", correctHash, downloadedHash)
	}
	err = Delete(fileName, false)
	if err != nil {
		t.Fatal(err)
	}
	err = Download(fileName)
	if err == nil {
		t.Fatal("file should not exist")
	}
}

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
