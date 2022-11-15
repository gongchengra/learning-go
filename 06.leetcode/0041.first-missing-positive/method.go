package main

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for k := range nums {
		if nums[k] != k+1 {
			return k + 1
		}
	}
	return len(nums) + 1
}
