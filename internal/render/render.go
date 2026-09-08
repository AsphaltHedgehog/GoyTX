package render

import "math"

func degreesToRadius(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
