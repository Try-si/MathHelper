package math

import "math"

func V2Normalize(V2 [2]float64) [2]float64 {
	magnitude := math.Sqrt(V2[0]*V2[0] + V2[1]*V2[1])
	return [2]float64{V2[0] / magnitude, V2[1] / magnitude}
}

func V2Distance(V2a [2]float64, V2b [2]float64) float64 {
	return math.Sqrt((V2a[0]-V2b[0])*(V2a[0]-V2b[0]) + (V2a[1]-V2b[1])*(V2a[1]-V2b[1]))
}
