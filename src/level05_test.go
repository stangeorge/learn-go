package main

import (
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Run("Blank filename check", func(t *testing.T) {
		err := createFile("")
		if err == nil {
			t.Error("Blank filename error not caught")
		}
	})

	t.Run("Create a file", func(t *testing.T) {
		f := "hello.txt"
		err := createFile(f)
		if err != nil {
			t.Errorf("Error creating file: %s", f)
		}

		file, err := os.Open(f)
		if err != nil {
			t.Error("Could not open file")
		}
		file.Close()
	})

	// t.Run("Verify file is created", func(t *testing.T) {
	// 	f := "hello"
	// 	createFile(f)
	// }
}
