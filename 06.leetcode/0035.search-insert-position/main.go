package main

import "fmt"

func main() {
	//     fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	//     fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	//     fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)>>1
		if nums[m] >= target {
			h = m - 1
		} else {
			if m == len(nums)-1 || nums[m+1] >= target {
				return m + 1
			}
			l = m + 1
		}
	}
	return 0
}
