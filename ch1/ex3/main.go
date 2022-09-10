package main

// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient version and the one that uses `strings.Join`.

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inefficientEchoTime := measureFunc(inefficientEcho)
	stringsJoinEchoTime := measureFunc(stringsJoinEcho)
	fmt.Println("inefficient echo time:", inefficientEchoTime)
	fmt.Println("strings join echo time:", stringsJoinEchoTime)
}

func inefficientEcho() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func stringsJoinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func measureFunc(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
