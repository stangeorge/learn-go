package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func level03() {
	const Max = 50000
	var n, r [Max]int
	sorts := []func([]int){selectionSort, bubbleSort}
	for _, s := range sorts {
		name := runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name()

		fmt.Printf("* %20s - Sorted  Ascending", name)
		for i := 0; i < Max; i++ {
			n[i] = i + 1
		}
		sortTime(s, n[:])

		fmt.Printf("* %20s - Sorted Descending", name)
		for i := 0; i < Max; i++ {
			r[i] = Max - i
		}
		sortTime(s, r[:])
	}
}

func sortTime(f func([]int), n []int) {
	// fmt.Println("\nn", n)
	fmt.Printf("(%d)", len(n))
	start := time.Now()
	f(n)
	stop := time.Now()
	// fmt.Println("\nn", n)
	fmt.Println(": ", stop.Sub(start))
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
