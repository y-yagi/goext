package osext

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsExit(t *testing.T) {
	pwd, _ := os.Getwd()

	file := filepath.Join(pwd, "osext_test.go")
	if !IsExist(file) {
		t.Errorf("Expect isExist returns true but false. file: %s", file)
	}

	file = filepath.Join(pwd, "unexist.go")
	if IsExist(file) {
		t.Errorf("Expect isExist returns false but true. file: %s", file)
	}
}
