package main

import (
	"flag"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"

	"code.google.com/p/draw2d/draw2d"
	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		nDots  = flag.Int("dots", 100, "number of dots")
		radius = flag.Int("radius", 10, "dot radius in px")
	)

	img, gc := graphicfarts.Setup(color.White)

	gc.SetLineWidth(3)
	gc.SetLineCap(draw2d.ButtCap)

	colors := []color.Color{
		color.RGBA{R: 119, G: 137, B: 121, A: 255},
		color.RGBA{R: 204, G: 220, B: 193, A: 255},
		color.RGBA{R: 240, G: 229, B: 209, A: 255},
		color.RGBA{R: 242, G: 218, B: 178, A: 255},
		color.RGBA{R: 215, G: 130, B: 88, A: 255},
	}

	for i := 0; i < *nDots; i++ {
		c := colors[i%len(colors)]
		gc.SetStrokeColor(c)
		gc.SetFillColor(c)

		x := rand.Intn(img.Rect.Dx())
		y := rand.Intn(img.Rect.Dy())

		gc.ArcTo(
			float64(x),
			float64(y),
			float64(*radius),
			float64(*radius),
			0,
			2*math.Pi,
		)
		gc.Fill()
		gc.Stroke()
	}

	if err := png.Encode(os.Stdout, img); err != nil {
		panic(err)
	}
}
