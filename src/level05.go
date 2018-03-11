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

func countPrefix(dir string, f string, n int) (s string, i int, err error) {
	var m map[string]int
	var prefix = ""
	var max = 0

	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		fin, err := os.Open(dir + f)
		if err != nil {
			err = &FileError{err.Error(), f}
		}
		defer fin.Close()

		scanner := bufio.NewScanner(fin)
		m = make(map[string]int)
		for scanner.Scan() {
			p := scanner.Text()[0:n]
			m[p] = m[p] + 1
			if m[p] > max {
				max = m[p]
				prefix = p
			}
		}
	}
	return prefix, max, err
}

func groupPrefix(dir string, f string, n int) (m map[string]int, err error) {
	if len(f) == 0 {
		err = ErrFileNameRequired
	} else {
		fin, err := os.Open(dir + f)
		if err != nil {
			err = &FileError{err.Error(), f}
		}
		defer fin.Close()

		scanner := bufio.NewScanner(fin)
		m = make(map[string]int)
		for scanner.Scan() {
			p := scanner.Text()[0:n]
			m[p] = m[p] + 1
		}
	}
	return m, err
}

//https://groups.google.com/d/msg/golang-nuts/FT7cjmcL7gw/Gj4_aEsE_IsJ
func sortPrefix(dir string, f string, n int) (err error) {
	var _, e = groupPrefix(dir, f, n)
	return e
}
