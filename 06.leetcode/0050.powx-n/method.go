package main

func myPow(x float64, n int) float64 {
	m := map[int]float64{}
	return power(x, n, m)
}

func power(x float64, n int, m map[int]float64) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 || x == 1 {
		return 1
	}
	if val, ok := m[n]; ok {
		return val
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		p := myPow(x, n/2)
		m[n/2] = p
		return p * p
	} else {
		p := myPow(x, (n-1)/2)
		m[(n-1)/2] = p
		return x * p * p
	}
}
