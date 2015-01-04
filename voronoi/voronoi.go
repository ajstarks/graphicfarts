package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		nSites = flag.Int("sites", 10, "number of sites")
	)

	canvas, rect := graphicfarts.Setup()

	var sx, sy []int
	var sc []string
	for i := 0; i < *nSites; i++ {
		x := rand.Intn(rect.Dx())
		y := rand.Intn(rect.Dy())

		r := uint8(rand.Intn(256))
		g := uint8(rand.Intn(256))
		b := uint8(rand.Intn(256))

		sx = append(sx, x)
		sy = append(sy, y)
		sc = append(sc, fmt.Sprintf("fill:none;stroke:#%02x%02x%02x", r, g, b))
	}

	// BUG(coda): holy shit this is the wrong way to do this

	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			dMin := dot(rect.Dx(), rect.Dy())
			var sMin int
			for s := 0; s < *nSites; s++ {
				if d := dot(sx[s]-x, sy[s]-y); d < dMin {
					sMin = s
					dMin = d
				}
			}
			canvas.Rect(x, y, 1, 1, sc[sMin])
		}
	}

	canvas.End()
}

func dot(x, y int) int {
	return x*x + y*y
}
