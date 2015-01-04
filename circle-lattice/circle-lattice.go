package main

import (
	"flag"
	"image/color"
	"image/png"
	"math"
	"os"

	"code.google.com/p/draw2d/draw2d"
	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		radius = flag.Int("radius", 50, "radius of circles in px")
		space  = flag.Int("space", 50, "spacing of circles in px")
	)

	img, gc := graphicfarts.Setup(color.White)

	gc.SetLineWidth(3)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetStrokeColor(color.Black)

	for y := img.Rect.Min.Y; y < img.Rect.Max.Y+*radius; y += *radius + *space {
		for x := img.Rect.Min.X; x < img.Rect.Max.X+*radius; x += *radius + *space {
			gc.ArcTo(
				float64(x),
				float64(y),
				float64(*radius),
				float64(*radius),
				0,
				2*math.Pi,
			)
			gc.Stroke()
		}
	}

	if err := png.Encode(os.Stdout, img); err != nil {
		panic(err)
	}
}
