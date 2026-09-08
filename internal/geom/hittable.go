package geom

import "github.com/AsphaltHedgehog/GoyTX/internal/vec"

type HitRecord struct {
	P      vec.Point3
	Normal vec.Vec3
	T      float64

	FrontFace bool
}

type Hittable interface {
	Hit(r Ray, rayT Interval) (HitRecord, bool)
}

func (h HitRecord) setFaceNormal(r Ray, outwardNormal vec.Vec3) {
	// Outward normal is assumed to have unit length.

	h.FrontFace = r.Dir.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Neg()
	}
}
