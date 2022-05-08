package restapi

import (
	"os"
	"path/filepath"

	"github.com/loqutus/O-1/pkg/file"
	"github.com/loqutus/O-1/pkg/types"
)

// MoveFile moves file from one directory to another.
func MoveFile(fileName string) error {
	file.EnsureDir(filepath.Join(types.Server.LocalDir, filepath.Dir(fileName)))
	err := os.Rename(fileName, filepath.Join(types.Server.LocalDir, fileName))
	return err
}
