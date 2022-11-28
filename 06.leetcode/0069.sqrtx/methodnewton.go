package main

func mySqrt(x float64) float64 {
	limit, res := 0.000001, x
	for abs(res-x/res) > limit {
		res = (x/res + res) / 2.0
	}
	return res
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
