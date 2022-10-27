package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			min = i
		}
	}
	first := binarySearch(nums[min:], target)
	if first != -1 {
		return min + first
	}
	last := binarySearch(nums[:min], target)
	if last != -1 {
		return last
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := (l + h) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			h = m - 1
		} else {
			return m
		}
	}
	return -1
}
