package main

// Exercise 1.1: Modify the `echo` program to also print `os.Args[0]`, the
// name of the command that invoked it.

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
