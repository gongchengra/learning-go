package main

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 || x == 1 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
