package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func level03() {
	logger.SetOutput(os.Stdout)
	const Max = 50000
	var n, r [Max]int
	sorts := []func([]int){selectionSort, bubbleSort, insertionSort}
	for _, s := range sorts {
		name := runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name()

		logger.Printf("* %20s - Sorted  Ascending", name)
		for i := 0; i < Max; i++ {
			n[i] = i + 1
		}
		sortTime(s, n[:])

		logger.Printf("* %20s - Sorted Descending", name)
		for i := 0; i < Max; i++ {
			r[i] = Max - i
		}
		sortTime(s, r[:])
	}
}

func sortTime(f func([]int), n []int) {
	// logger.Printf("(%d) %v BEFORE", len(n), n)
	start := time.Now()
	f(n)
	stop := time.Now()
	// logger.Printf("(%d) %v AFTER", len(n), n)
	logger.Println("Time to sort: ", stop.Sub(start))
}

func selectionSort(n []int) {
	for i := 0; i < len(n); i++ {
		min_j, min := i, n[i]
		for j := i; j < len(n); j++ {
			if n[j] < min {
				min_j, min = j, n[j]
			}
		}
		if min_j != i {
			n[i], n[min_j] = n[min_j], n[i]
		}
	}
}

func bubbleSort(n []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				swapped = true
				n[i], n[i+1] = n[i+1], n[i]
			}
		}
	}
}

func insertionSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		if n[i+1] < n[i] {
			for j := i + 1; j < len(n)-1; j++ {
				if n[j] < n[j-1] {
					n[j-1], n[j] = n[j], n[j-1]
					// logger.Printf("i=%d, j=%d, n[i]=%d, n[j]=%d\n",
					// 	i, j, n[i], n[j])
					break
				}
			}
		}
	}
}
