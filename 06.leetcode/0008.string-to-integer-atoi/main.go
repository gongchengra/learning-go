package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
}

func myAtoi(str string) int {
	flag := 1
	s := 0
	start := 0
	for _, v := range str {
		if v == ' ' {
			if start == 0 {
				continue
			} else {
				break
			}
		} else if v == '+' {
			if start == 0 {
				start = 1
				flag = 1
			} else {
				break
			}
		} else if v == '-' {
			if start == 0 {
				start = 1
				flag = -1
			} else {
				break
			}
		} else if v < '0' || v > '9' {
			if start == 0 {
				return 0
			} else {
				break
			}
		} else {
			start = 1
			s = 10*s + int(v-'0')
			if flag == 1 && s > math.MaxInt32 {
				return math.MaxInt32
			}
			if flag == -1 && -s < math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	res := 0
	if flag == 1 {
		res = s
	} else {
		res = -s
	}
	return res
}
