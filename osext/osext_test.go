package osext

import (
	"os"
	"testing"
)

func TestIsExit(t *testing.T) {
	pwd, _ := os.Getwd()

	file := pwd + "/osext_test.go"
	if !IsExist(file) {
		t.Errorf("Expect isExist returns true but false. file: %s", file)
	}

	file = pwd + "/unexist.go"
	if IsExist(file) {
		t.Errorf("Expect isExist returns false but true. file: %s", file)
	}
}
