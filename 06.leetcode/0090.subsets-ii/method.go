package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var dfs func(int, []int)
	dfs = func(idx int, tmp []int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)
		for i := idx; i < len(nums); i++ {
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(tmp, nums[i]))
			}
		}
	}
	dfs(0, make([]int, 0, len(nums)))
	return res
}
