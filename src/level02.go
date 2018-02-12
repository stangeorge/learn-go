package main

import (
	"fmt"
	"runtime"
	"time"
)

func level02() {
	fmt.Print("\n* Fibonacci Using Recursion")
	timeTaken(fibonacciRecursive, int64(40))

	fmt.Print("\n* Fibonacci Using Iteration")
	timeTaken(fibonacciIterative, int64(40))

	var n int64 = 2000000000
	fmt.Println("\n* 4 times in a row: Fibonacci Using Iteration")
	fourTimes := func(n int64) int64 {
		for i := 0; i < 4; i++ {
			timeTaken(fibonacciIterative, n)
		}
		return 0
	}
	timeTaken(fourTimes, n)

	fmt.Println("\n* 4 times in a row: Fibonacci Using Iteration And Concurrency")
	start := time.Now()
	c := make(chan int64)
	for i := 0; i < 4; i++ {
		go fibonacciIterativeConcurrent(n, c)
		fmt.Println(": ", <-c)
	}
	stop := time.Now()
	fmt.Println(": ", stop.Sub(start))

	fmt.Println("\n* 4 times in a row: Fibonacci Using Iteration And Concurrency With Multiple CPUs: ",
		runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) //number of CPUs
	start = time.Now()
	c = make(chan int64)
	for i := 0; i < 4; i++ {
		go fibonacciIterativeConcurrent(n, c)
		fmt.Println(": ", <-c)
	}
	stop = time.Now()
	fmt.Println(": ", stop.Sub(start))
}

func timeTaken(f func(int64) int64, i int64) {
	start := time.Now()
	fmt.Print(": ", f(i))
	stop := time.Now()
	fmt.Println(": ", stop.Sub(start))
}

func fibonacciRecursive(n int64) int64 {
	switch n {
	case 0:
		return 0
	case 1:
		fallthrough
	case 2:
		return 1
	default:
		return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	}
}

func fibonacciIterative(n int64) int64 {
	var x, y int64 = 0, 1
	for i := n; i > 1; i-- {
		x, y = y, x+y
	}
	return y
}

func fibonacciIterativeConcurrent(n int64, c chan int64) {
	c <- fibonacciIterative(n)
}
