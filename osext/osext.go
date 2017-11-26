package osext

import "os"

// IsExist checks the file exists.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
