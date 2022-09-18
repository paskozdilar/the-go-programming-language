package main

// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact, and assess
// its per formance.

import (
	"fmt"
	"math/rand"
	"time"
)

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

// Rewritten popcount to use a bitshift and decrement
func PopCountShiftDecr(x uint64) int {
	y := 0
	for x != 0 {
		y += int(x & 1)
		x = (x - 1) & x
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
	MeasureFunc(PopCountShiftDecr)
	fmt.Println("PopCount:", MeasureFunc(PopCount))
	fmt.Println("PopCountShift:", MeasureFunc(PopCountShiftDecr))
}
