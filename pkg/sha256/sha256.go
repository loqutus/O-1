package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Sha256 calculates the SHA256 hash of a file.
func GetFileSHA256(filePath string) (string, error) {
	logrus.Println("GetFileSHA256 " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	buf := make([]byte, 1024*1024)
	for {
		bytesRead, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		hash.Write(buf[:bytesRead])
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
