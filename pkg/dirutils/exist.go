package dirutils

import (
	"os"
)

// Exist checks if the given filepath exist in the
// cwd, if any err occur when trying to read the info
// then the error is returned
func Exist(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return false, nil
	}
	return true, nil
}
