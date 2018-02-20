package main

import (
	"os"
	"reflect"
	"runtime"
	"testing"
)

const Max = 50000

var n, r, sorted, reverse_sorted [Max]int
var sorts = []func([]int){selectionSort, bubbleSort, insertionSort}

func initAscending(a []int) {
	for i := 0; i < Max; i++ {
		a[i] = i + 1
	}
}

func initDescending(a []int) {
	for i := 0; i < Max; i++ {
		a[i] = Max - i
	}
}

func TestSorting(t *testing.T) {
	logger.SetOutput(os.Stdout)
	initAscending(sorted[:])
	initDescending(reverse_sorted[:])

	for _, sort := range sorts {
		name := runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name()
		t.Run("Sorted List", func(t *testing.T) {
			initAscending(n[:])
			if n != sorted {
				t.Errorf("%s failed to initialize a sorted list", name)
			}
			sort(n[:])
			if n != sorted {
				t.Errorf("%s failed to sort a sorted list", name)
			}
		})

		t.Run("Reverse Sorted List", func(t *testing.T) {
			initDescending(r[:])
			if r != reverse_sorted {
				t.Errorf("%s failed to initialize a reverse sorted list", name)
			}
			// logger.Printf("BEFORE:%v (%s)\n", r, name)
			sort(r[:])
			// logger.Printf("AFTER :%v (%s)\n", r, name)
			if r != sorted {
				t.Errorf("%s failed to sort a reverse sorted list", name)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	// logger.SetOutput(os.Stdout)
	initAscending(sorted[:])
	initDescending(reverse_sorted[:])
	initAscending(n[:])
	initDescending(r[:])
	b.ResetTimer()

	b.Run("Sorted List", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			selectionSort(n[:])
		}
	})

	b.Run("Reverse Sorted List", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			selectionSort(r[:])
		}
	})
}
