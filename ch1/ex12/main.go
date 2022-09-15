package main

// Exercise 1.12: Modify the Lissajous server to read parameter values from the
// URL. For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of
// the default 5. Use the strconv.Atoi function to convert the string parameter
// into an integer. You can see its documentation with go doc strconv.Atoi.

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(404)
			return
		}
		p := Parameters{
			cycles:  5,      // number of complete x oscillator revolutions
			res:     0.0001, // angular resolution
			size:    100,    // image canvas covers [-size..+size]
			nframes: 64,     // number of animation frames
			delay:   8,      // delay between frames in 10ms units
		}
		cycles, err := strconv.Atoi(r.Form.Get("cycles"))
		if err == nil {
			p.cycles = cycles
		}
		res, err := strconv.ParseFloat(r.Form.Get("res"), 64)
		if err == nil {
			p.res = res
		}
		size, err := strconv.Atoi(r.Form.Get("size"))
		if err == nil {
			p.size = size
		}
		nframes, err := strconv.Atoi(r.Form.Get("nframes"))
		if err == nil {
			p.nframes = nframes
		}
		delay, err := strconv.Atoi(r.Form.Get("delay"))
		if err == nil {
			p.delay = delay
		}
		lissajous(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Parameters struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func lissajous(out io.Writer, p Parameters) {
	freq := rand.Float64() * 3.0 // relative frequency of the y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles*2)*math.Pi; t += p.res {
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
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5), colorIndex)
		}
		phase += 0.01
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
