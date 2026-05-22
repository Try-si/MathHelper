package math

import "math"

func V2Normalize(V2 [2]float64) [2]float64 {
	magnitude := math.Sqrt(V2[0]*V2[0] + V2[1]*V2[1])
	return [2]float64{V2[0] / magnitude, V2[1] / magnitude}
}
