// Genertae gif animations of random Lissajous figures
// run as ./lissajous.exe > out.gif
package main

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
	color.Black, color.White, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	blackIndex = 0
	whiteIndex = 1
	greenIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // no. x oscillations
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size...+size]
		nframes = 64    // no. animation frames
		delay   = 8     // no. ms between frames
	)
	freq := rand.Float64() * 3.0 // relative freq. of oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorInd := i%2 + 1
		for t := 0.0; t < 2*cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				uint8(colorInd))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
