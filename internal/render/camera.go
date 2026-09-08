package render

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/AsphaltHedgehog/GoyTX/internal/geom"
	"github.com/AsphaltHedgehog/GoyTX/internal/ppm"
	"github.com/AsphaltHedgehog/GoyTX/internal/vec"
)

type Options struct {
	Seed     uint64
	Progress io.Writer
}

type Camera struct {
	AspectRatio float64 // Ratio of image width over height
	ImageWidth  int     // Rendered image width in pixel count

	imageHeight int        // Rendered image height
	center      vec.Point3 // Camera center
	pixel00Loc  vec.Point3 // Location of pixel 0, 0
	pixelDeltaU vec.Vec3   // Offset to the pixel to the right
	pixelDeltaV vec.Vec3   // Offset to the pixel below
}

func NewCamera(aspectRatio float64, imgWidth int) Camera {
	return Camera{
		AspectRatio: aspectRatio,
		ImageWidth:  imgWidth,
	}
}

func (c Camera) Render(world geom.HittableList) {
	c.initialize()

	pw, err := ppm.NewWriter(os.Stdout, c.ImageWidth, c.imageHeight)
	if err != nil {
		panic(err)
	}

	opts := Options{
		Seed:     0,
		Progress: os.Stderr,
	}

	for j := range c.imageHeight {
		fmt.Fprintf(opts.Progress, "\rScanlines remaining: %d ", c.imageHeight-j)
		for i := range c.ImageWidth {
			pixelDeltaUI := c.pixelDeltaU.Scale(float64(i))
			pixelDeltaVJ := c.pixelDeltaV.Scale(float64(j))
			pixelCenter := c.pixel00Loc.Add(pixelDeltaUI).Add(pixelDeltaVJ)

			rayDirection := pixelCenter.Sub(c.center)
			r := geom.Ray{Orig: c.center, Dir: rayDirection}

			pixelColor := c.rayColor(r, &world)

			err := pw.WritePixel(pixelColor)

			if err != nil {
				panic(err)
			}

		}
	}

	if err := pw.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stderr, "\rDone.                 \n")
}

func (c *Camera) initialize() {
	c.imageHeight = max(int(float64(c.ImageWidth)/c.AspectRatio), 1)
	c.center = vec.Point3{}

	// Viewpoint Dimension
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	// Vector calc across the horizontal and down the vertical viewport edges.
	viewportU := vec.Vec3{X: viewportWidth}
	viewportV := vec.Vec3{Y: -viewportHeight}

	// Calculate pixel deltas
	c.pixelDeltaU = viewportU.Div(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.Div(float64(c.imageHeight))

	focal := vec.Vec3{Z: focalLength}
	halfU := viewportU.Div(2)
	halfV := viewportV.Div(2)
	viewPortUpperLeft := c.center.
		Sub(focal).
		Sub(halfU).
		Sub(halfV)

	halfPixelDeltaSum := c.pixelDeltaU.Add(c.pixelDeltaV).Scale(0.5)
	c.pixel00Loc = viewPortUpperLeft.Add(halfPixelDeltaSum)
}

func (c Camera) rayColor(r geom.Ray, world *geom.HittableList) vec.Color {
	if rec, ok := world.Hit(r, geom.Interval{Min: 0, Max: math.Inf(1)}); ok {
		return rec.Normal.Add(vec.Color{X: 1, Y: 1, Z: 1}).Scale(0.5)
	}

	unitDir := r.Dir.Unit()
	a := 0.5 * (unitDir.Y + 1.0)

	return vec.Color{X: 1, Y: 1, Z: 1}.Scale(1.0 - a).Add(vec.Color{X: 0.5, Y: 0.7, Z: 1.0}.Scale(a))
}
