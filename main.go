package main

import (
	"github.com/AsphaltHedgehog/GoyTX/internal/geom"
	"github.com/AsphaltHedgehog/GoyTX/internal/render"
	"github.com/AsphaltHedgehog/GoyTX/internal/vec"
)

func main() {
	// World setup
	world := geom.HittableList{}
	world.Add(geom.Sphere{Center: vec.Point3{Z: -1}, Radius: 0.5})
	world.Add(geom.Sphere{Center: vec.Point3{Y: -100.5, Z: -1}, Radius: 100})

	// Image parameters setup
	aspectRatio := 16.0 / 10.0
	imgWidth := 400
	camera := render.NewCamera(aspectRatio, imgWidth)

	camera.Render(world)
}
