package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(a, 2))
	fmt.Println(a)
}

func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
