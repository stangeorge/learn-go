package main

import (
	"fmt"
	"time"
)

func level02() {
	fmt.Print("\n* Fibonacci Using Recursion")
	timeTaken(fibonacciRecursive, 40)

	fmt.Print("\n* Fibonacci Using Iteration")
	timeTaken(fibonacciIterative, 40)
}

func timeTaken(f func(int) int, i int) {
	start := time.Now()
	fmt.Print(": ", f(i))
	stop := time.Now()
	fmt.Println(": ", stop.Sub(start))
}

func fibonacciRecursive(n int) int {
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

func fibonacciIterative(n int) int {
	x, y := 0, 1
	for i := n; i > 1; i-- {
		x, y = y, x+y
	}
	return y
}
