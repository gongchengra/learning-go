package main

import "sort"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	help(nums, 0, &res)
	return res
}

func help(nums []int, idx int, res *[][]int) {
	if idx == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	sort.Ints(nums[idx:])
	for i := idx; i < len(nums); i++ {
		if i != idx && nums[i-1] == nums[i] {
			continue
		}
		nums[i], nums[idx] = nums[idx], nums[i]
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		help(tmp, idx+1, res)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
