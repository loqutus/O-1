package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func GetFileSHA256(filePath string) (string, error) {
	logrus.Println("GetFileSHA256 " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
