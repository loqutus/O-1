package sha256

import (
	"crypto/md5"
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

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
