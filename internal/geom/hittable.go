package geom

import "github.com/AsphaltHedgehog/GoyTX/internal/vec"

type HitList struct {
	P      vec.Point3
	Normal vec.Vec3
	t      float64
}
type Hittable interface {
	hit(r Ray, rayTMin float64, rayTMax float64, hitRecord HitList) bool
}
