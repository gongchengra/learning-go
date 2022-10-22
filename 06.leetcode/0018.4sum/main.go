package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[3] <= target; i++ {
		if (i > 0 && nums[i] == nums[i-1]) ||
			nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if (j > i+1 && nums[j] == nums[j-1]) ||
				nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				if sum := nums[i] + nums[j] + nums[l] + nums[r]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
