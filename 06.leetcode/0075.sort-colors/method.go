package main

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		if nums[j] < 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j+1
		} else if nums[j] > 1 {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}
}
