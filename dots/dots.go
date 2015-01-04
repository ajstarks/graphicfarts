package main

import (
	"flag"
	"math/rand"

	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		nDots  = flag.Int("dots", 100, "number of dots")
		radius = flag.Int("radius", 10, "dot radius in px")
	)

	canvas, rect := graphicfarts.Setup()

	colors := []string{
		"#778979",
		"#ccdcc1",
		"#f0e5d1",
		"#f2dab2",
		"#d78258",
	}

	for i := 0; i < *nDots; i++ {
		c := colors[i%len(colors)]

		x := rand.Intn(rect.Dx())
		y := rand.Intn(rect.Dy())

		canvas.Circle(x, y, *radius, "stroke:none;fill:"+c)
	}

	canvas.End()
}
