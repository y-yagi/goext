package osext

import "os"

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
