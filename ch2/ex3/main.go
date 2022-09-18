package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
// Compare the performance of the two versions. (Section 11.4 shows how to
// compare the performance of different implementations systematically.)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Rewritten popcount to use a loop
func PopCountLoop(x uint64) int {
	y := 0
	for i := 0; i < 8; i++ {
		y += int(pc[byte(x>>(i*8))])
	}
	return y
}

func MeasureFunc(fn func(x uint64) int) time.Duration {
	x := rand.Uint64()
	start := time.Now()
	fn(x)
	return time.Since(start)
}

func main() {
	// Call functions once to nullify startup time
	MeasureFunc(PopCount)
	MeasureFunc(PopCountLoop)
	fmt.Println("PopCount:", MeasureFunc(PopCount))
	fmt.Println("PopCountLoop:", MeasureFunc(PopCountLoop))
}
