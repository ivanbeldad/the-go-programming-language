// Modify the Lissajous program to produce images in multiple colors by adding more
// values to palette and then displaying them by changing the third argument of Set-ColorIndex
// in some interesting way.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var backgroundColor = color.Black

const (
	numberOfColors = 10
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func createPalette(numberOfColors int) []color.Color {
	palette := make([]color.Color, numberOfColors+1)
	palette[0] = color.Black
	for i := 1; i < numberOfColors+1; i++ {
		newColor := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
		palette[i] = newColor
	}
	return palette
}

func lissajous(out io.Writer) {
	palette := createPalette(numberOfColors)
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// randomColor := uint8(rand.Intn(len(palette)-1) + 1)
			strInt := string(strconv.Itoa(time.Now().Nanosecond())[1])
			randomColor, _ := strconv.Atoi(strInt)
			// fmt.Println(strconv.Itoa(int(randomColor)))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randomColor+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
