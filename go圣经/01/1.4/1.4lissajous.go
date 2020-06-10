// Lissajous generates GIF animations of random Lissajous figures.
package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
)

var palette = []color.Color{color.White, color.Black}

type Lconfig struct {
	Cycles float64
	Res    float64
	Freq   float64
	Size   int
	Nframes int
	Delay  int
}


const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func Lissajous(out io.Writer ,  lconf Lconfig) {

	//freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: lconf.Nframes}
	phase := 0.0 // phase difference
	for i := 0; i < lconf.Nframes; i++ {
		rect := image.Rect(0, 0, 2*lconf.Size+1, 2*lconf.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < lconf.Cycles*2*math.Pi; t += lconf.Res {
			x := math.Sin(t)
			y := math.Sin(t*lconf.Freq + phase)
			img.SetColorIndex(lconf.Size+int(x*float64(lconf.Size)+0.5), lconf.Size+int(y*float64(lconf.Size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, lconf.Delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if err != nil {fmt.Println(err)}
}