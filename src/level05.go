package main

import (
	"bufio"
	"bytes"
	"errors"
	"html/template"
	"os"
	"os/exec"
	"sort"
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

//Pair :A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

//PairList :A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

//https://groups.google.com/d/msg/golang-nuts/FT7cjmcL7gw/Gj4_aEsE_IsJ
func sortPrefix(dir string, f string, n int) (PairList, error) {
	var m, err = groupPrefix(dir, f, n)

	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p, err
}

func findMeanMedian(dir string, f string, n int) (mean float64, median float64, err error) {
	p, err := sortPrefix(dir, f, 5)
	if err != nil {
		logger.Printf("Unable to sort")
	}
	//Make a slice out of p.Value
	sum := 0
	values := make([]int, 0, len(p))
	for _, v := range p {
		values = append(values, v.Value)
		sum += v.Value
	}

	l := len(values)
	mean = float64(sum) / float64(l)

	if l%2 == 0 {
		median = float64(values[l/2])
	} else {
		median = float64(values[l/2-1]+values[l/2-1]) / 2
	}
	return mean, median, err
}

func textTemplate(s string, p []Pair) (string, error) {
	tmpl := s + `{{range .}}{{.Value}}, {{end}}`
	b := new(bytes.Buffer)
	t, err := template.New("").Parse(tmpl)
	err = t.Execute(b, p)
	return b.String(), err
}
