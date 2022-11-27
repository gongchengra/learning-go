package main

func plusOne(digits []int) []int {
	flag := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+flag > 9 {
			digits[i] = 0
			flag = 1
		} else {
			digits[i] += flag
			flag = 0
		}
	}
	if flag > 0 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}
