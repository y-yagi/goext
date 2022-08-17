package ioext_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/y-yagi/goext/ioext"
)

func TestCopyFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "ioexttest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	tmpfn1 := filepath.Join(tempDir, "tmpfile1")
	content := []byte("temporary file's content")
	if err := os.WriteFile(tmpfn1, content, 0644); err != nil {
		t.Fatal(err)
	}

	tmpfn2 := filepath.Join(tempDir, "tmpfile2")
	if err := ioext.CopyFile(tmpfn1, tmpfn2, 0644); err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(tmpfn2)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(content) {
		t.Fatalf("got '%q', want '%q'", got, content)
	}
}
