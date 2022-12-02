package main

func removeDuplicates(nums []int) (res int) {
	for fast, v := range nums {
		if fast < 2 || nums[res-2] != v {
			nums[res] = v
			res++
		}
	}
	return
}
