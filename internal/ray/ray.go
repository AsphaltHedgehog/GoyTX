package ray

import "goytx/m/internal/vec"

type Ray struct {
	Orig vec.Point3
	Dir  vec.Vec3
}

func (r Ray) At(t float64) vec.Point3 {
	return r.Orig.Add(r.Dir.Scale(t))
}
