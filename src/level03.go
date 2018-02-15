package main

import (
	"fmt"
	"time"
)

func level03() {
	const Max = 10
	var n, r [Max]int
	for i := 0; i < Max; i++ {
		n[i] = i + 1
	}
	// fmt.Println("\nn", n)

	for i := 0; i < Max; i++ {
		r[i] = Max - i
	}
	// fmt.Println("\nr", r)

	fmt.Print("\n* Selection Sort - Sorted List")
	sortTime(selectionSort, n[:])
	fmt.Print("\n* Selection Sort - Reverse Sorted List")
	sortTime(selectionSort, r[:])
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
