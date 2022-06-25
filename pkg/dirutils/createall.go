package dirutils

import (
	"os"
	"path/filepath"
)

func CreateAll(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.FileMode(0755)); err != nil {
		return nil, err
	}
	return os.Create(path)
}
