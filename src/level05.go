package main

import (
	"os/exec"
)

func level05() {
}

//FileError to handle errors in file operatoins
type FileError struct {
	Message string
	File    string
}

//FileNameRequiredError to handle blank file names
func FileNameRequiredError(message string, file string) *FileError {
	return &FileError{
		Message: message,
		File:    file,
	}
}

//FileCreateError to handle error during file creation
func FileCreateError(message string, file string) *FileError {
	return &FileError{
		Message: message,
		File:    file,
	}
}

func (e *FileError) Error() string {
	return e.Message + " " + e.File
}

// https://golang.org/pkg/os/exec/#example_Command
func createFile(f string) (err error) {
	if len(f) == 0 {
		err = FileNameRequiredError("file name required", f)
	}
	cmd := exec.Command("touch", f)
	err = cmd.Run()
	if err != nil {
		logger.Fatal(err)
		return FileCreateError("could not create file", f)
	}
	return err
}
