package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Exercise 2.2: Write a general-purpose unit-conversion program analogous to
// cf that reads numbers from its command-line arguments or from the standard
// input if there are no arguments, and converts each number into units like
// temperature in Celsius and Fahrenheit, length in feet and meters, weight
// in pounds and kilograms, and the like.

func main() {
	for arg := range Input() {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		c := Celsius(n)
		f := Fahrenheit(n)
		ft := Feet(n)
		m := Meter(n)
		lb := Pound(n)
		kg := Kilogram(n)
		km := Kilometer(n)
		mi := Mile(n)
		fmt.Printf("%s = %s, %s = %s; %s = %s, %s = %s; %s = %s, %s = %s; %s = %s, %s = %s\n",
			c, c.CToF(), f, f.FToC(),
			ft, ft.FtToM(), m, m.MToFt(),
			lb, lb.LbToKg(), kg, kg.KgToLb(),
			km, km.KmToMi(), mi, mi.MiToKm(),
		)
	}
}

func Input() chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		if len(os.Args[1:]) > 0 {
			for _, arg := range os.Args[1:] {
				ch <- arg
			}
		} else {
			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				ch <- s.Text()
			}
		}
	}()
	return ch
}

type (
	Celsius    float64
	Fahrenheit float64
	Feet       float64
	Meter      float64
	Pound      float64
	Kilogram   float64
	Kilometer  float64
	Mile       float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%.3g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3g°F", f) }
func (ft Feet) String() string      { return fmt.Sprintf("%.3gft", ft) }
func (m Meter) String() string      { return fmt.Sprintf("%.3gm", m) }
func (lb Pound) String() string     { return fmt.Sprintf("%.3glb", lb) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%.3gkg", kg) }
func (km Kilometer) String() string { return fmt.Sprintf("%.3gkm", km) }
func (mi Mile) String() string      { return fmt.Sprintf("%.3gmi", mi) }

func (c Celsius) CToF() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (f Fahrenheit) FToC() Celsius { return Celsius((f - 32) * 5 / 9) }
func (ft Feet) FtToM() Meter       { return Meter(ft / 3.2808) }
func (m Meter) MToFt() Feet        { return Feet(m * 3.2808) }
func (lb Pound) LbToKg() Kilogram  { return Kilogram(lb / 2.2046) }
func (kg Kilogram) KgToLb() Pound  { return Pound(kg * 2.2046) }
func (km Kilometer) KmToMi() Mile  { return Mile(km * 0.62) }
func (mi Mile) MiToKm() Kilometer  { return Kilometer(mi / 0.62) }
