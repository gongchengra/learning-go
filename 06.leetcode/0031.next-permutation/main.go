package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, k := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums)
	}
	for k = len(nums) - 1; i >= 0 && k > i; k-- {
		if nums[i] < nums[k] {
			break
		}
	}
	if i >= 0 {
		nums[i], nums[k] = nums[k], nums[i]
		s := nums[i+1 : len(nums)]
		reverse(s)
		return
	}
}

func reverse(nums []int) {
	for i, l := 0, len(nums); i < l-i; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
