# Learn Go
Learning the Go language by building over simple examples **(Work In Progress)**  

### SETUP
Installing go on macOS
```
brew update
brew upgrade
brew install go --cross-compile-common
```

To run the examples here:
```
git clone https://github.com/stangeorge/learn-go.git
go run learn-go/src/*.go
```

Change prompt to just $ using: `export PS1="$ "`

### LEVEL 1
Source Code [/src/level01.go](/src/level01.go)  
[Find Environment Variables](#find-environment-variables)  
[Get the variables separately using range](#get-the-variables-separately-using-range)  
[Get both index and values variables separately](#get-both-index-and-values-variables-separately)  
[Ignore index using underscore](#ignore-index-using-underscore)  
[Split the values](#split-the-values)  
[Get only the environment variable names](#get-only-the-environment-variable-names)  
[Call a function](#call-a-function)  
[Concurrency](#concurrency)  
[Have the routines finish before main](#have-the-routines-finish-before-main)  
[Parallel run on multiple CPUs](#parallel-run-on-multiple-cpus)  

### LEVEL 2
Source Code [/src/level02.go](/src/level02.go)  
[Fibonacci Using Recursion](#fibonacci-using-recursion)  
[Fibonacci Using Iteration](#fibonacci-using-iteration)  
[Passing a function as an argument](#passing-a-function-as-an-argument)  
[Anonymous Functions](#anonymous-functions)  
[4 times in a row: Fibonacci Using Iteration](#4-times-in-a-row-fibonacci-using-iteration)  
[4 times in a row: Fibonacci Using Iteration And Concurrency](#4-times-in-a-row-fibonacci-using-iteration-and-concurrency)  
[4 times in a row: Fibonacci Using Iteration And Concurrency With Multiple CPUs](#4-times-in-a-row-fibonacci-using-iteration-and-concurrency-with-multiple-cpus)  

### Find Environment Variables
**Concepts:** package, imports, main function, printing a line, running a go program.

    package main

    import (
      "fmt"
      "os"
    )

    func main() {
      fmt.Println(os.Environ())
    }

>$ go run learn.go  
>[TERM_PROGRAM=vscode VIRTUALENVWRAPPER_SCRIPT=/usr/local/bin/virtualenvwrapper.sh VIRTUALENVWRAPPER_PROJECT_FILENAME=.project TERM...]

**Results:** I got all the variables in a single block. Let me try to separate this out.

___

### Get the variables separately using range
**Concepts:** creating an initializing a variable, for-loop, range

    for e := range os.Environ() {
      fmt.Println(e)
    }

>$ go run learn.go  
>0  
>1  

**Results:** I got only the index values. Let me get the variables.

___

### Get both index and values variables separately
**Concepts:** indexes and values in a for-loop

    for i, e := range os.Environ() {
      fmt.Println(i, e)
    }

>$ go run learn.go  
>0 TERM_PROGRAM=vscode  
>1 VIRTUALENVWRAPPER_SCRIPT=/usr/local/bin/virtualenvwrapper.sh  

**Results:** I got the index and variables. Let me ignore the index.

___

### Ignore index using underscore
**Concepts:** underscore

    for _, e := range os.Environ() {
      fmt.Println(e)
    }

>$ go run learn.go  
>TERM_PROGRAM=vscode  
>VIRTUALENVWRAPPER_SCRIPT=/usr/local/bin/virtualenvwrapper.sh  

**Results:**I got the variable-value pair. Let me split them out.

___

### Split the values
**Concepts:** strings.Split(), array

    for _, e := range os.Environ() {
      pair := strings.Split(e, "=")
      fmt.Println(pair)
    }

>$ go run learn.go  
>[TERM_PROGRAM vscode]  
>[VIRTUALENVWRAPPER_SCRIPT /usr/local/bin/virtualenvwrapper.sh]  

**Results:** I split the values and got arrays

___

### Get only the environment variable names
**Concepts:** array

    for _, e := range os.Environ() {
      pair := strings.Split(e, "=")
      fmt.Println(pair[0])
    }
    
>$ go run learn.go  
>TERM_PROGRAM  
>VIRTUALENVWRAPPER_SCRIPT  

**Results:** I got the first element in an array

___

### Call a function
**Concepts:** function, parameters

    func printEnviron(e string) {
        time.Sleep(10 * time.Millisecond)
        fmt.Println(e)
    }

    func main() {
        for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            printEnviron(pair[0])
        }
    }

>$ time go run learn.go  
>  
>real    0m0.616s  
>user    0m0.146s  
>sys     0m0.097s  

**Results:** Called a function in a loop. Let me see if I can reduce the execution time

___

### Concurrency
**Concepts:** go routines, channels

    func printEnviron(e string, channel chan string) {
        time.Sleep(10 * time.Millisecond)
        fmt.Println(e)
        channel <- e
    }

    func main() {
        channel := make(chan string)
        for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            go printEnviron(pair[0], channel)
        }
        fmt.Println(<-channel) //outside the for-loop
    }

$ time go run learn.go  
>VIRTUALENVWRAPPER_VIRTUALENV  
>  
>real    0m0.291s  
>user    0m0.219s  
>sys     0m0.103s  

**Results:** Called a method concurrently many times. They communicate with the main thread using channel and pass it the variable name. Real time for this was 0m0.291s vs 0m0.616s if this was called sequentially in a loop. HOWEVER, this does not mean that all the routines finished before the main thread finished.

___

### Have the routines finish before main
**Concepts:** go routines, channels

    func main() {
        channel := make(chan string)
        for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            go printEnviron(pair[0], channel)
            fmt.Println(<-channel) //inside the for-loop
        }
    }

$ time go run learn.go  
>TERM_PROGRAM 
>  
>real    0m0.654s  
>user    0m0.205s  
>sys     0m0.109s  

**Results:** Moving the `fmt.Println(<-channel)` inside the for-loop makes the main function wait till it gets a response from all the routines. Note that this took around the same 0.654s as the prior sequential run that took 0.616s. Lets see if we can make this run in parallel.

___

### Parallel run on multiple CPUs
**Concepts:** runtime, cores

    func main() {
        runtime.GOMAXPROCS(runtime.NumCPU()) //number of CPUs
        
        channel := make(chan string)
        for _, e := range os.Environ() {
            pair := strings.Split(e, "=")
            go printEnviron(pair[0], channel)
            fmt.Println(<-channel)
        }
    }

$ time go run learn.go  
>TERM_PROGRAM  
>  
>real    0m0.646s  
>user    0m0.206s  
sys     0m0.109s  

**Results:** We set `GOMAXPROCS` to the max CPUs we have. This did not seem to affect the execution time. It was still around 0.646s vs the sequential 0.616s. So no gain in speed yet.

___

### Fibonacci Using Recursion
**Concepts:** Recursion, switch-case.

    func fibonacciRecursive(n int) int {
        switch n {
        case 0:
            return 0
        case 1:
            fallthrough
        case 2:
            return 1
        default:
        default:
            return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
        }
    }
    func main() {
    	fmt.Println(fibonacciRecursive(50))
    }

>$ time go run learn.go  
>12586269025  

>real    2m36.339s  
>user    2m31.408s  
>sys     0m1.105s   

**Results:** Recursion in action. However, it takes 2 minutes 36 seconds to find the 50th number in the Fibonacci series. Lets see if there is a faster way.

___

### Fibonacci Using Iteration
**Concepts:** for-loop

    func fibonacciIterative(n int) int {
        x, y := 0, 1
        for i:=n; i > 1; i-- {
            x, y = y, x+y
        }
        return y
    }
    func main() {
	    fmt.Println(fibonacciRecursive(50))
    }
    
>$time go run learn.go  
>12586269025  
>  
>real    0m0.276s  
>user    0m0.204s  
>sys     0m0.102s  
    
**Results:** Faster execution. It takes only 276ms vs 2m36s in the recursive version

___

### Passing a function as an argument
Say I want to find the time taken by the two Fibonacci functions above. I could start and stop the timer before calling each function or I could create a function to do it. `timeTaken(f func(int) int, i int)` takes a function as an argument. The function we pass it takes in an integer hence `func(int)`. It also returns an integer, hence `func(int) int`

    func timeTaken(f func(int) int, i int) {
        start := time.Now()
        fmt.Print(": ", f(i))
        stop := time.Now()
        fmt.Println(": ", stop.Sub(start))
    }
    func main() {
        fmt.Print("\n* Fibonacci Using Recursion")
        timeTaken(fibonacciRecursive, 40)

        fmt.Print("\n* Fibonacci Using Iteration")
        timeTaken(fibonacciIterative, 40)
    }

>  
>  * Fibonacci Using Recursion: 102334155:  634.742717ms  
>  
>  * Fibonacci Using Iteration: 102334155:  2.036µs  

___

### Anonymous Functions
I need to run the function 4 times in a row and print the timing. I won't be reusing this, so I don't really need a new function. So I can create an anonymous function and use it.

    fmt.Print("\n* 4 times in a row: Fibonacci Using Recursion\n")
	fourTimes := func(n int) int {
		for i := 0; i < 4; i++ {
			timeTaken(fibonacciRecursive, n)
		}
		return 0
	}
	timeTaken(fourTimes, 40)

> * 4 times in a row: Fibonacci Using Recursion  
>: 102334155:  599.31207ms  
>: 102334155:  659.258468ms  
>: 102334155:  631.379041ms  
>: 102334155:  627.767538ms  
>: 0:  2.517787938s  

___

### 4 times in a row: Fibonacci Using Iteration

    var n int64 = 2000000000
	fmt.Println("\n* 4 times in a row: Fibonacci Using Iteration")
	fourTimes := func(n int64) int64 {
		for i := 0; i < 4; i++ {
			timeTaken(fibonacciIterative, n)
		}
		return 0
	}
	timeTaken(fourTimes, n)

> * 4 times in a row: Fibonacci Using Iteration  
>: 2697763845588227525:  1.099461051s  
>: 2697763845588227525:  1.102832038s  
>: 2697763845588227525:  1.12779254s  
>: 2697763845588227525:  1.137225242s  
>: 0:  4.467368851s  

___

### 4 times in a row: Fibonacci Using Iteration And Concurrency
    start := time.Now()
	c := make(chan int64)
	for i := 0; i < 4; i++ {
		go fibonacciIterativeConcurrent(n, c)
		fmt.Println(": ", <-c)
	}
	stop := time.Now()
	fmt.Println(": ", stop.Sub(start))

> * 4 times in a row: Fibonacci Using Iteration And Concurrency
>:  2697763845588227525  
>:  2697763845588227525  
>:  2697763845588227525  
>:  2697763845588227525  
>:  4.438340355s  

___

### 4 times in a row: Fibonacci Using Iteration And Concurrency With Multiple CPUs

    runtime.GOMAXPROCS(runtime.NumCPU()) //number of CPUs
	start = time.Now()
	c = make(chan int64)
	for i := 0; i < 4; i++ {
		go fibonacciIterativeConcurrent(n, c)
		fmt.Println(": ", <-c)
	}
	stop = time.Now()
	fmt.Println(": ", stop.Sub(start))

> * 4 times in a row: Fibonacci Using Iteration And Concurrency With Multiple CPUs:  4  
>:  2697763845588227525  
>:  2697763845588227525  
>:  2697763845588227525  
>:  2697763845588227525  
>:  4.429574803s  

___

### Analysis
I've annotated the spikes due to the code in red. Notice that the most of the activity is on cores 1 and 3 based on the spikes of green squares. There is barely any action on cores 2 and 4 corresponding to the annotations. My theory is that although we set GOMAXPROCS to the number of CPUs, the go scheduler works differently.  

![Activity in CPU Cores](/fib_annotated.png)  

Also see these responses from the FAQ on golang.org:  
[Why doesn't my multi-goroutine program use multiple CPUs? ¶](https://golang.org/doc/faq#Why_no_multi_CPU)  
[Why does using GOMAXPROCS > 1 sometimes make my program slower?](https://golang.org/doc/faq#Why_GOMAXPROCS)

___

### LEVEL 3
Source Code [/src/level02.go](/src/level03.go)  
[Sorting Setup](#sorting-setup)  
[Selection Sort](#selection-sort)  

### Sorting Setup
I created 2 arrays. "n" has numbers from 1 to 50,000 in sorted order. "r" has the numbers reverse sorted. I have a function `sortTime` to measure the time taken

    func main() {
        const Max = 50000
        var n, r [Max]int
        for i := 0; i < Max; i++ {
            n[i] = i + 1
        }

        for i := 0; i < Max; i++ {
            r[i] = Max - i
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

___

### Selection Sort

    func selectionSort(n []int) {
        for i := 0; i < len(n); i++ {
            min_j, min := i, n[i]
            for j := i; j < len(n); j++ {
                if n[j] < min {
                    min_j, min = j, n[j]
                }
            }
            n[i], n[min_j] = n[min_j], n[i]
        }
    }
    
    func main() {
        fmt.Print("\n* Selection Sort - Sorted List")
        sortTime(selectionSort, n[:])
        fmt.Print("\n* Selection Sort - Reverse Sorted List")
        sortTime(selectionSort, r[:])
    }

Output with 50,000 numbers
> * Selection Sort - Sorted List:  1.907869346s  
> * Selection Sort - Reverse Sorted List:  1.675308716s  

Output with 10 numbers
> * Selection Sort - Sorted List(10):  1.03µs  
> * Selection Sort - Reverse Sorted List(10):  778ns  

Oddly, the sorted list takes longer than the reverse sorted list. I tried to add a check to avoid unnecessary swaps but that did not help either:

    if (min_j != i) {
        n[i], n[min_j] = n[min_j], n[i]
    }

___
