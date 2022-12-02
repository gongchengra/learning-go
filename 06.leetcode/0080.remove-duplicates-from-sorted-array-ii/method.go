package main

func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 2 {
		return ln
	}
	res := 2
	for i := 2; i < ln; i++ {
		if nums[i] != nums[res-2] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
