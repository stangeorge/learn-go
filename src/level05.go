package main

import (
	"errors"
	"os"
	"os/exec"
)

var (
	//ErrFileNameRequired to handle blank file names
	ErrFileNameRequired = errors.New("file name is required")
)

func level05() {
}

//FileError to handle errors in file operations
type FileError struct {
	Message string
	File    string
}

func (e *FileError) Error() string {
	return e.Message + " " + e.File
}

func createFile(f string) (err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		_, err = os.Create(f)
		if err != nil {
			err = &FileError{err.Error(), f}
		}
	}
	return err
}

func extractFile(f string) (err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		cmd := exec.Command("7z", "x", f)
		err = cmd.Run()
		if err != nil {
			err = &FileError{err.Error(), f}
		}
	}
	return err
}
