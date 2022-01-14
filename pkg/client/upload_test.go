package client

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"

	fileSHA256 "github.com/loqutus/O-1/pkg/sha256"
	"github.com/loqutus/O-1/pkg/types"
)

func TestUpload(t *testing.T) {
	types.Client.HostName = "localhost"
	types.Client.Port = "6969"
	file, err := os.CreateTemp(".", "o-1-test*")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())
	data := []byte("abc abc abc")
	if _, err := file.Write(data); err != nil {
		t.Fatal(err)
	}
	fileName := filepath.Base(file.Name())
	if err := Upload(fileName); err != nil {
		t.Fatal(err)
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		t.Fatal(err)
	}
	correctHash := hex.EncodeToString(hash.Sum(nil))
	err = Download(fileName)
	if err != nil {
		t.Fatal(err)
	}
	downloadedHash, err := fileSHA256.GetFileSHA256(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if correctHash != downloadedHash {
		t.Fatalf("hash mismatch: %s != %s", correctHash, downloadedHash)
	}
}
