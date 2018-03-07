package main

import (
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Run("Blank filename check", func(t *testing.T) {
		f := ""
		err := createFile(f)
		if err != ErrFileNameRequired {
			t.Error("Blank filename error not caught")
		}
	})

	t.Run("Create a file", func(t *testing.T) {
		var f = "/.txt"
		err := createFile(f)
		if err != err.(*FileError) {
			t.Errorf("Invalid filename error not caught")
		}
	})

	t.Run("Extract a 7z file", func(t *testing.T) {
		var f = os.Getenv("HOME") + "/Downloads/pwned-passwords-update-2.txt.7z"
		err := extractFile(f)
		if err != nil {
			t.Errorf("Error extracting file %s: %s", f, err.Error())
		}
	})
}
