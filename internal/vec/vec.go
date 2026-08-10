package vec

type Vec3 struct {
	X, Y, Z float64
}

type Point3 = Vec3
type Color = Vec3

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec3) Mul(u Vec3) Vec3 {
	return Vec3{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func (v Vec3) Div(t float64) Vec3 {
	return v.Scale(1 / t)
}

func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) Length() float64 {
	return v.LengthSquared()
}

func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

func (v Vec3) Unit() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{(u.Y*v.Z - u.Z*v.Y), (u.Z*v.X - u.X*v.Z), (u.X*v.Y - u.Y*v.X)}
}
