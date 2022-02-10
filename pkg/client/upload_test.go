package client

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	fileSHA256 "github.com/loqutus/O-1/pkg/sha256"
	"github.com/loqutus/O-1/pkg/types"
)

func TestUpload(t *testing.T) {
	for i := 0; i < 1000; i++ {
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
	if err := Upload(fileName, false); err != nil {
		t.Fatal(err)
	}
	correctHash, err := fileSHA256.GetFileSHA256(fileName)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
	os.Remove(file.Name())
	err = Download(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fileName)
	downloadedHash, err := fileSHA256.GetFileSHA256(fileName)
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
