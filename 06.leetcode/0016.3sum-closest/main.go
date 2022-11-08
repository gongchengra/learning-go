package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
