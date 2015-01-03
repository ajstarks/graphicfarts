package graphicfarts

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	"code.google.com/p/draw2d/draw2d"
)

var (
	width  = flag.Int("width", 500, "width in px")
	height = flag.Int("height", 500, "height in px")
	seed   = flag.Int64("seed", 0, "seed for PRNG (-1 for random)")
)

func Setup(bg color.Color) (*image.RGBA, *draw2d.ImageGraphicContext) {
	flag.Parse()
	randomize(*seed)

	img := image.NewRGBA(image.Rect(0, 0, *width, *height))
	if bg != nil {
		draw.Draw(img, img.Rect, &image.Uniform{C: bg}, image.Pt(0, 0), draw.Over)
	}

	gc := draw2d.NewGraphicContext(img)

	return img, gc
}

func randomize(i int64) {
	if i == -1 {
		if err := binary.Read(crand.Reader, binary.LittleEndian, &i); err != nil {
			panic(err)
		}
	}
	rand.Seed(i)
}
