package main

// Exercise 3.4: Following the approach of the Lissajous example in Section
// 1.7, construct a web server that computes surfaces and writes SVG data to
// the client. The server must set the Content-Type header like this:
//
//		w.Header().Set("Content-Type", "image/svg+xml")
//
// (This step was not required in the Lissajous example because the server uses
// standard heuristics to recognize common formats like PNG from the first 512
// bytes of the response and generates the proper header.) Allow the client to
// specify values like height, width, and color as HTTP request parameters.

// NOTE: For some reason, SVG won't render in neither Firefox nor Chrome.
// TODO: Figure out why.

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	color         = "#000000"           // color of polygon
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(404)
			return
		}
		p := Parameters{width, height, color}
		width, err := strconv.Atoi(r.Form.Get("width"))
		if err == nil {
			p.width = width
		}
		height, err := strconv.Atoi(r.Form.Get("height"))
		if err == nil {
			p.height = height
		}
		color := r.Form.Get("color")
		if len(color) == 7 {
			_, err = strconv.Atoi(color[1:])
			if color[0] == '#' && err == nil {
				p.color = color
			}
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(200)
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j, p)
				bx, by := corner(i, j, p)
				cx, cy := corner(i, j+1, p)
				dx, dy := corner(i+1, j+1, p)
				if anyInvalidFloat(ax, ay, bx, by, cx, cy, dx, dy) {
					// skip polygons with infinite value
					continue
				}
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type Parameters struct {
	width  int
	height int
	color  string
}

func corner(i, j int, p Parameters) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(p.width)/2 + (x-y)*cos30*xyscale
	sy := float64(p.height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func anyInvalidFloat(vals ...float64) bool {
	isInf := false
	for _, val := range vals {
		if math.IsInf(val, 0) || math.IsNaN(val) {
			isInf = true
		}
	}
	return isInf
}
