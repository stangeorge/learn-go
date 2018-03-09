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
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt.7z"

		err := extractFile(dir, f)
		if err != nil {
			t.Errorf("Error extracting file %s: %s", dir+f, err.Error())
		}

	})

	t.Run("Count lines in a file", func(t *testing.T) {
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt"
		n, err := countLines(dir, f)
		if err != nil {
			t.Errorf("Error extracting file %s: %s", dir+f, err.Error())
		}
		count := 399790

		if n != count {
			t.Errorf("Expected %d lines but counted only %d", count, n)
		}

		//Cleanup the extracted file
		os.Remove(dir + f)
	})
}
