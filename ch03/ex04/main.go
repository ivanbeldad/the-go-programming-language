// Following the approach of the Lissajous example in Section 1.7, construct a web
// server that computes surfaces and writes SVG dat a to the client.

package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 20.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/svg", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "image/svg+xml")
		writeSvg(res)
	})
	http.ListenAndServe("localhost:8000", nil)
}

func writeSvg(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			fill := getColorByHeight(i, j)
			r, g, b, _ := fill.RGBA()
			ax, ay, okA := corner(i+1, j)
			bx, by, okB := corner(i, j)
			cx, cy, okC := corner(i, j+1)
			dx, dy, okD := corner(i+1, j+1)
			if !okA || !okB || !okC || !okD {
				continue
			}
			fmt.Fprintf(writer, "<polygon style='fill:#%2.2X%2.2X%2.2X' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				uint8(r), uint8(g), uint8(b), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(writer, "</svg>")
}

func corner(i, j int) (sx, sy float64, ok bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

func f(x, y float64) (f float64, ok bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	f = math.Sin(r) / r
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0, false
	}
	return math.Sin(r) / r, true
}

func getColorByHeight(i, j int) color.Color {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, _ := f(x, y)
	// fmt.Printf("height: %f\n", z)

	// HH HH HH -> 00000000 00000000 00000000
	// -1 -> 0 -> 1
	// -1 = 00000000 00000000 11111111
	// 0 = 00000000 11111111 00000000
	// 1 = 11111111 00000000 00000000
	// total shift 16
	shift := uint8(math.Round((z + 1) * 8))
	var customColor uint32 = 0xFF
	customColor = customColor << shift
	b := uint8(customColor)
	customColor = customColor >> 8
	g := uint8(customColor)
	customColor = customColor >> 8
	r := uint8(customColor)
	// fmt.Printf("Color: %b\tR: %b\tG: %b\tB: %b\n", originalColor, r, g, b)
	// fmt.Printf("Shift: %d\n", shift)
	return color.RGBA{r, g, b, 0xFF}
}
