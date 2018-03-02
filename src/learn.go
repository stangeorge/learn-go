package main

import (
	"bytes"
	"log"
	"os"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	logger.SetOutput(os.Stdout)
	// logger.Println("* LEVEL 1")
	// level01()
	// logger.Println("* LEVEL 2")
	// level02()
	// logger.Println("* LEVEL 3")
	// level03()
	logger.Println("* LEVEL 5")
	level05()
}
