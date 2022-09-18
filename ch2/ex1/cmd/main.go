package main

import (
	"fmt"

	tempconv "github.com/paskozdilar/the-go-programming-language/ch2/ex1"
)

func main() {
	c := tempconv.Celsius(0)
	f := tempconv.Fahrenheit(0)
	k := tempconv.Kelvin(0)

	fmt.Println(c, "=", tempconv.CToF(c), "=", tempconv.CToK(c))
	fmt.Println(f, "=", tempconv.FToC(f), "=", tempconv.FToK(f))
	fmt.Println(k, "=", tempconv.KToC(k), "=", tempconv.KToF(k))
}
