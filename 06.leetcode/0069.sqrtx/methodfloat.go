package main

func mySqrt(x float64) float64 {
	left, right, limit := 0.0, x/2, 0.000001
	for left <= right {
		mid := left + (right-left)/2
		if abs(mid*mid-x) < limit {
			return mid
		} else if mid*mid < x {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
