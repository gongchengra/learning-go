package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, abs(nums[0]+nums[1]+nums[2]-target)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < target {
				l++
				if delta > target-s {
					delta = target - s
					res = s
				}
			} else if s == target {
				return s
			} else {
				r--
				if delta > s-target {
					delta = s - target
					res = s
				}
			}
		}
	}
	return res
}
