// Generating GIF animations.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // full rounds
		res     = 0.001 // angle
		size    = 100   // size of image [-size.. +size]
		nframes = 64    // how many frames of animation
		delay   = 8     // delay between frames (10ms)
	)
	rand.Seed(time.Now().UnixNano())

	freq := rand.Float64() * 3.0 // freq of osc y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	random := rand.Intn(5-1) + 1
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+1), size+int(y*size+1), uint8(random))

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim) // Note: ignoring coding errors
}
