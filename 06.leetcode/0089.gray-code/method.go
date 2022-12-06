package main

func grayCode(n int) []int {
	var l uint = 1 << uint(n)
	res := make([]int, l)
	for i := uint(0); i < l; i++ {
		res[i] = int((i >> 1) ^ i)
	}
	return res
}
