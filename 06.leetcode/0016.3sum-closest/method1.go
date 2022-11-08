package main

import "sort"

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
