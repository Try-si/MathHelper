package math

func V2Normilize(V2 [2]int) [2]int {
	return [2]int{V2[0] / (V2[0]*V2[0] + V2[1]*V2[1]), V2[1] / (V2[0]*V2[0] + V2[1]*V2[1])}
}
