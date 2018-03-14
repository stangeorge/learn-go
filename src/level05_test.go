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
		var f = "pwned-passwords-update-2.txt"

		//Cleanup the extracted file
		os.Remove(dir + f)

		err := extractFile(dir, f+".7z")
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
	})

	t.Run("Get the prefix that occurs the most number of time", func(t *testing.T) {
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt"

		p, n, err := countPrefix(dir, f, 5)
		if err != nil {
			t.Errorf("Error extracting file %s: %s", dir+f, err.Error())
		}

		expected := "36DC1"
		count := 6

		if n != count || p != expected {
			t.Errorf("Expected %s to occur %d times", p, n)
		}
	})

	t.Run("Group prefixes", func(t *testing.T) {
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt"

		m, err := groupPrefix(dir, f, 5)
		expected := 332324
		if err != nil || len(m) != expected {
			t.Errorf("Error grouping prefixes for file %s: %s", dir+f, err.Error())
		}

	})

	t.Run("Sort prefixes", func(t *testing.T) {
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt"

		p, err := sortPrefix(dir, f, 5)
		// expected := {"A1C8A", "36DC1", "F30EC"}
		check := p[0].Value == 6 && p[1].Value == 6 && p[2].Value == 6 && p[3].Value == 5

		if err != nil || !check {
			t.Errorf("Error grouping prefixes for file %s: %s", dir+f, err.Error())
		}

	})

	t.Run("Median", func(t *testing.T) {
		var dir = os.Getenv("HOME") + "/Downloads/"
		var f = "pwned-passwords-update-2.txt"

		mean, median, err := findMeanMedian(dir, f, 5)
		if err != nil || mean != 1.203012722523802 || median != 1 {
			t.Errorf("Error finding median in file %s: %s", dir+f, err.Error())
		}

	})
}
