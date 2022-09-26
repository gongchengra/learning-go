package Fibonacci

// Saving numbers into a slice

func FibonacciSequence(num int) []float64 {
	var res = []float64{}
	res = append(res, 0)
	if num == 0 {
		return res
	}
	res = append(res, 1)
	if num == 1 {
		return res
	}
	for n := 2; n <= num; n++ {
		result := res[n-1] + res[n-2]
		res = append(res, result)
	}
	return res
}
