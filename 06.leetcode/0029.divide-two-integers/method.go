package main

func divide(dividend int, divisor int) int {
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend = -dividend
	}
	if divisor < 0 {
		flag *= -1
		divisor = -divisor
	}
	res, sum := 1, divisor
	for {
		sum <<= 1
		res <<= 1
		if sum > dividend {
			break
		}
	}
	for {
		sum -= divisor
		res--
		if sum <= dividend {
			break
		}
	}
	if flag < 0 {
		return -res
	}
	if res > ((1 << 31) - 1) {
		return ((1 << 31) - 1)
	}
	return res
}
