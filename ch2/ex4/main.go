package main

// Exercise 2.4: Write a version of PopCount that counts bits by shifting its
// argument through 64 bit positions, testing the rightmost bit each time.
// Compare its performance to the table lookup version.

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

// Rewritten popcount to use a bitshift and compare
func PopCountShift(x uint64) int {
	y := 0
	for x != 0 {
		y += int(x & 1)
		x <<= 1
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
	MeasureFunc(PopCountShift)
	fmt.Println("PopCount:", MeasureFunc(PopCount))
	fmt.Println("PopCountShift:", MeasureFunc(PopCountShift))
}
