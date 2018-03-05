package main

import (
	"errors"
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

//FileCreateError to handle error during file creation
func FileCreateError(file string) *FileError {
	return &FileError{
		Message: "error creating file",
		File:    file,
	}
}

func (e *FileError) Error() string {
	return e.Message + " " + e.File
}

// https://golang.org/pkg/os/exec/#example_Command
func createFile(f string) (err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		cmd := exec.Command("touch", f)
		err = cmd.Run()
		if err != nil {
			// logger.Fatal(err)
			err = FileCreateError(f)
		}
	}
	return err
}
