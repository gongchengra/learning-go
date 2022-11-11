package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}

/*

func reverse(x int) int {
	ret := []int{}
	a := 0
	if x > 0 {
		a = x
	} else {
		a = -x
	}
	for q := a; q > 0; q = q / 10 {
		r := q % 10
		ret = append(ret, r)
	}
	s, res := 0, 0
	for _, v := range ret {
		s = 10*s + v
	}
	if x > 0 {
		res = s
	} else {
		res = -s
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res
}
*/

func reverse(x int) (res int) {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x = x / 10
	}
	res = sign * res
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return
}
