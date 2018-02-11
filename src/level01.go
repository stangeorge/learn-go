package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func level01() {
	fmt.Println("\n* Find Environment Variables")
	fmt.Println(os.Environ())

	fmt.Println("\n* Get the variables separately using range")
	for e := range os.Environ() {
		fmt.Println(e)
	}

	fmt.Println("\n* Get both index and values variables separately")
	for i, e := range os.Environ() {
		fmt.Println(i, e)
	}

	fmt.Println("\n* Ignore index using underscore")
	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	fmt.Println("\n* Split the values")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}

	fmt.Println("\n* Get only the environment variable names")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

	fmt.Println("\n* Call a function")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		printEnviron(pair[0])
	}

	fmt.Println("\n* Concurrency")
	c1 := make(chan string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		go printEnvironConcurrent(pair[0], c1)
	}
	fmt.Println(<-c1) //outside the for-loop

	fmt.Println("\n* Have the routines finish before main")
	c2 := make(chan string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		go printEnvironConcurrent(pair[0], c2)
		fmt.Println(<-c2) //inside the for-loop
	}

	fmt.Println("\n* ### Parallel run on multiple CPUs")
	runtime.GOMAXPROCS(runtime.NumCPU()) //number of CPUs
	c3 := make(chan string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		go printEnvironConcurrent(pair[0], c3)
		fmt.Println(<-c3)
	}
}

func printEnviron(e string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println(e)
}

func printEnvironConcurrent(e string, channel chan string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println(e)
	channel <- e
}
