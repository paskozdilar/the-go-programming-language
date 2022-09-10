package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exercise 1.4: Modify `dup2` to print the names of all files in which each
// duplicated line occurs.

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "/dev/stdin", filenames)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, file, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			for filename, _ := range filenames[line] {
				fmt.Printf("\t%s", filename)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, filenames map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == nil {
			filenames[input.Text()] = make(map[string]bool)
		}
		filenames[input.Text()][filename] = true
	}
}
