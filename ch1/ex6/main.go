package main

// Exercise 1.6: Modify the Lissajous program to produce images in multiple
// colors by adding more values to palette and then displaying them by changing
// the third argument of SetColorIndex in some interesting way.

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.White,
}

const (
	blackIndex = 0 // first color in palette
	redIndex   = 1 // second color in palette
	greenIndex = 2 // third color in palette
	blueIndex  = 3 // fourth color in palette
	whiteIndex = 4 // fifth color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5      // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 100    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of the y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Depending on which quadrant current point is in, use one of the
			// four non-black colors.
			colorIndex := uint8(1)
			if x > 0 {
				colorIndex += 1
			}
			if y > 0 {
				colorIndex += 2
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.01
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
