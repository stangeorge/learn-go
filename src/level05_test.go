package main

import (
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
		var f = "`.txt"
		err := createFile(f)
		if err != nil {
			t.Errorf("Error creating file: %s", f)
		}
	})
}
