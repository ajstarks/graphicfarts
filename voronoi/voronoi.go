package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/codahale/graphicfarts"
	"github.com/pzsz/voronoi"
)

func main() {
	var (
		nSites = flag.Int("sites", 10, "number of sites")
	)

	canvas, rect := graphicfarts.Setup()
	rand.Seed(int64(time.Now().Nanosecond()) % 1e9)
	var sites []voronoi.Vertex
	for i := 0; i < *nSites; i++ {
		x := rand.Intn(rect.Dx())
		y := rand.Intn(rect.Dy())

		sites = append(sites, voronoi.Vertex{X: float64(x), Y: float64(y)})
	}

	bbox := voronoi.NewBBox(0, float64(rect.Max.X), 0, float64(rect.Max.Y))
	diagram := voronoi.ComputeDiagram(sites, bbox, true)

	canvas.Gstyle("stroke:black")
	for _, cell := range diagram.Cells {
		var x, y []int
		for _, halfedge := range cell.Halfedges {
			start := halfedge.GetStartpoint()
			x = append(x, int(start.X))
			y = append(y, int(start.Y))

			end := halfedge.GetEndpoint()
			x = append(x, int(end.X))
			y = append(y, int(end.Y))
		}

		x = append(x, x[0])
		y = append(y, y[0])

		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)

		style := fmt.Sprintf("fill:#%02x%02x%02x", r, g, b)
		canvas.Polygon(x, y, style)
	}
	canvas.Gend()

	for _, site := range sites {
		canvas.Circle(int(site.X), int(site.Y), 5)
	}

	canvas.End()
}
