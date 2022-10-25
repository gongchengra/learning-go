package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

/*
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l-1] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
*/

func removeDuplicates(nums []int) int {
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[l] == nums[r] {
			continue
		}
		l++
		nums[l], nums[r] = nums[r], nums[l]
	}
	return l + 1
}
