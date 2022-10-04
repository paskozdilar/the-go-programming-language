package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// Exercise 3.5: Implement a full-color Mandelbrot set using the function
// image.NewRGBA and the type color.RGBA or color.YCbCr.

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 50

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			var red, green, blue uint8
			ratio := float64(n) / float64(iterations)
			value := uint8(ratio * 255)
			if value < 128 {
				blue = (127 - value) * 2
				green = value * 2
			} else {
				green = (255 - value) * 2
				blue = (value - 127) * 2
			}
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}
