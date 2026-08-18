package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AsphaltHedgehog/GoyTX/internal/geom"
	"github.com/AsphaltHedgehog/GoyTX/internal/ppm"
	"github.com/AsphaltHedgehog/GoyTX/internal/vec"
)

type Options struct {
	Seed     uint64
	Progress io.Writer
}

func rayColor(r geom.Ray, world geom.Hittable) vec.Color {
	if rec, ok := world.Hit(r, 0, math.Inf(1)); ok {
		return rec.Normal.Add(vec.Color{X: 1, Y: 1, Z: 1}).Scale(0.5)
	}

	unitDir := r.Dir.Unit()
	a := 0.5 * (unitDir.Y + 1.0)

	return vec.Color{X: 1, Y: 1, Z: 1}.Scale(1.0 - a).Add(vec.Color{X: 0.5, Y: 0.7, Z: 1.0}.Scale(a))
}

func main() {
	// Image parameters setup
	aspectRatio := 16.0 / 10.0
	imgWidth := 400
	imgHeight := int(float64(imgWidth) / aspectRatio)

	if imgHeight < 1 {
		imgHeight = 1
	}

	// World setup
	world := geom.HittableList{}

	world.Add(geom.Sphere{Center: vec.Point3{Z: -1}, Radius: 0.5})
	world.Add(geom.Sphere{Center: vec.Point3{Y: -100.5, Z: -1}, Radius: 100})

	// Camera and Viewport init
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imgWidth) / float64(imgHeight))
	cameraCenter := vec.Point3{}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vec.Vec3{X: viewportWidth}
	viewportV := vec.Vec3{Y: -viewportHeight}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Div(float64(imgWidth))
	pixelDeltaV := viewportV.Div(float64(imgHeight))

	// Calculate the location of the upper left pixel.
	focal := vec.Vec3{Z: focalLength}
	halfU := viewportU.Div(2)
	halfV := viewportV.Div(2)
	viewPortUpperLeft := cameraCenter.
		Sub(focal).
		Sub(halfU).
		Sub(halfV)

	halfPixelDeltaSum := pixelDeltaU.Add(pixelDeltaV).Scale(0.5)
	pixel00Loc := viewPortUpperLeft.Add(halfPixelDeltaSum)

	pw, err := ppm.NewWriter(os.Stdout, imgWidth, imgHeight)
	if err != nil {
		panic(err)
	}

	opts := Options{
		Seed:     0,
		Progress: os.Stderr,
	}

	// Render

	for j := range imgHeight {
		fmt.Fprintf(opts.Progress, "\rScanlines remaining: %d ", imgHeight-j)
		for i := range imgWidth {
			pixelDeltaUI := pixelDeltaU.Scale(float64(i))
			pixelDeltaVJ := pixelDeltaV.Scale(float64(j))
			pixelCenter := pixel00Loc.Add(pixelDeltaUI).Add(pixelDeltaVJ)

			rayDirection := pixelCenter.Sub(cameraCenter)
			r := geom.Ray{Orig: cameraCenter, Dir: rayDirection}

			pixelColor := rayColor(r, &world)

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
