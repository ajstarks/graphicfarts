package graphicfarts

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"image"
	"math/rand"
	"os"

	"github.com/ajstarks/svgo"
)

var (
	width  = flag.Int("width", 500, "width in px")
	height = flag.Int("height", 500, "height in px")
	seed   = flag.Int64("seed", 0, "seed for PRNG (-1 for random)")
)

func Setup(styles ...string) (*svg.SVG, image.Rectangle) {
	flag.Parse()
	randomize(*seed)

	canvas := svg.New(os.Stdout)
	canvas.Start(*width, *height)
	if styles != nil {
		canvas.Rect(0, 0, *width, *height, styles...)
	}

	return canvas, image.Rect(0, 0, *width, *height)
}

func randomize(i int64) {
	if i == -1 {
		if err := binary.Read(crand.Reader, binary.LittleEndian, &i); err != nil {
			panic(err)
		}
	}
	rand.Seed(i)
}
