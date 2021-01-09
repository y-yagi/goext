package osext

import (
	"io"
	"os"
)

// IsExist checks the file exists.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// IsEmptyDir checks a directory empty or not.
func IsEmptyDir(path string) (bool, error) {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}

	defer f.Close()

	if _, err = f.Readdirnames(1); err == io.EOF {
		return true, nil
	}

	return false, err
}
