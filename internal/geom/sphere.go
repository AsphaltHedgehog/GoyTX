package geom

import (
	"math"

	"github.com/AsphaltHedgehog/GoyTX/internal/vec"
)

type Sphere struct {
	Center vec.Point3
	Radius float64
}

func New(center vec.Point3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	oc := s.Center.Sub(r.Orig)
	a := r.Dir.LengthSquared()
	h := r.Dir.Dot(oc)
	c := oc.LengthSquared() - (s.Radius * s.Radius)
	discriminant := h*h - a*c

	if discriminant < 0 {
		return HitRecord{}, false
	}

	discSqrt := math.Sqrt(discriminant)

	root := (h - discSqrt) / a
	if root <= tMin || tMax <= root {
		root = (h + discSqrt) / a
		if root <= tMin || tMax <= root {
			return HitRecord{}, false
		}
	}

	rT := r.At(root)
	outNormal := rT.Sub(s.Center).Div(s.Radius)

	hitRec := HitRecord{
		T:      root,
		P:      rT,
		Normal: outNormal,
	}

	hitRec.setFaceNormal(r, outNormal)

	return hitRec, true
}
