package main

import (
	"bufio"
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

func extractFile(dir string, f string) (err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		cmd := exec.Command("7z", "x", dir+f, "-o"+dir)
		err = cmd.Run()
		if err != nil {
			err = &FileError{err.Error(), f}
		}
	}
	return err
}

func countLines(dir string, f string) (i int, err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		fin, err := os.Open(dir + f)
		if err != nil {
			err = &FileError{err.Error(), f}
		}
		defer fin.Close()

		scanner := bufio.NewScanner(fin)
		for scanner.Scan() {
			i++
		}
	}
	return i, err
}
