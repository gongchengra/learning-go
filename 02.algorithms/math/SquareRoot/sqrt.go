package SquareRoot

import "math"

func SquareRoot(num float64) float64 {
	// Method only supports positiv integers
	if num < 0 {
		return -1
	}
	// Handle edge case square root of zero
	if num == 0 {
		return 0
	}
	// Set start values
	root := 1.0
	cur := 0.0
	// Find root
	//for cur-root != 0 {
	for math.Abs(cur-root) > 0.00000001 {
		root = newton(root, num)
		cur = newton(root, num)
	}
	return root
}

func newton(z, x float64) float64 {
	return z - (((z * z) - x) / (2 * z))
}
