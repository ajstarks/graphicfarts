package main

import (
	"flag"

	"github.com/codahale/graphicfarts"
)

func main() {
	var (
		radius = flag.Int("radius", 50, "radius of circles in px")
		space  = flag.Int("space", 50, "spacing of circles in px")
	)

	canvas, rect := graphicfarts.Setup("fill:white")

	for y := rect.Min.Y; y < rect.Max.Y+*radius; y += *radius + *space {
		for x := rect.Min.X; x < rect.Max.X+*radius; x += *radius + *space {
			canvas.Gstyle("fill:none;stroke:black;stroke-width:3")
			canvas.Circle(x, y, *radius)
			canvas.Gend()
		}
	}

	canvas.End()
}
