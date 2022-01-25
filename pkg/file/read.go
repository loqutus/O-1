package file

import (
	"os"

	"github.com/loqutus/O-1/pkg/types"
)

func Read(fileName string) ([]byte, error) {
	fileBody, err := os.ReadFile(types.Server.LocalDir + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return fileBody, nil
}
