package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	sum, max, i := nums[0], nums[0], 1
	for i < len(nums) {
		fmt.Println(i, sum, nums[i])
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		fmt.Println(i, sum, nums[i])
		if max < sum {
			max = sum
		}
		i++
	}
	return max
}
