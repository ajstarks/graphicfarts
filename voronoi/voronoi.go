package main

import (
	"flag"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"code.google.com/p/draw2d/draw2d"
	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		nSites = flag.Int("sites", 10, "number of sites")
	)

	img, gc := graphicfarts.Setup(color.White)

	gc.SetLineWidth(3)
	gc.SetLineCap(draw2d.ButtCap)

	var sx, sy []int
	var sc []color.Color
	for i := 0; i < *nSites; i++ {
		x := rand.Intn(img.Rect.Dx())
		y := rand.Intn(img.Rect.Dy())
		c := color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: uint8(rand.Intn(256)),
		}
		sx = append(sx, x)
		sy = append(sy, y)
		sc = append(sc, c)
	}

	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			dMin := dot(img.Rect.Dx(), img.Rect.Dy())
			var sMin int
			for s := 0; s < *nSites; s++ {
				if d := dot(sx[s]-x, sy[s]-y); d < dMin {
					sMin = s
					dMin = d
				}
			}
			img.Set(x, y, sc[sMin])
		}
	}

	if err := png.Encode(os.Stdout, img); err != nil {
		panic(err)
	}
}

func dot(x, y int) int {
	return x*x + y*y
}
