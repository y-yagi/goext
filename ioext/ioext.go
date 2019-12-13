package ioext

import (
	"io"
	"os"
)

// CopyFile copies file
func CopyFile(src, dst string, perm os.FileMode) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	_, err = io.Copy(df, sf)
	err2 := df.Close()
	if err != nil {
		return err
	}
	return err2
}
