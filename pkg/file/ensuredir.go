package file

import (
	"os"

	"github.com/sirupsen/logrus"
)

// EnsureDir recursively creates a directory if it doesn't exist.
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		logrus.Println("Creating directory:", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
