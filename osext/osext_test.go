package osext_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/y-yagi/goext/osext"
)

func TestIsExit(t *testing.T) {
	pwd, _ := os.Getwd()

	file := filepath.Join(pwd, "osext_test.go")
	if !osext.IsExist(file) {
		t.Errorf("Expect isExist returns true but false. file: %s", file)
	}

	file = filepath.Join(pwd, "unexist.go")
	if osext.IsExist(file) {
		t.Errorf("Expect isExist returns false but true. file: %s", file)
	}
}

func TestIsEmptyDir(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "osexttest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	result, err := osext.IsEmptyDir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Errorf("Expect IsEmptyDir returns true but false. dir: %s", tempDir)
	}

	if err = ioutil.WriteFile(tempDir+"/hello", []byte("message"), 0644); err != nil {
		t.Fatal(err)
	}
	if result, err = osext.IsEmptyDir(tempDir); err != nil {
		t.Fatal(err)
	}
	if result {
		t.Errorf("Expect IsEmptyDir returns false but true. dir: %s", tempDir)
	}
}
