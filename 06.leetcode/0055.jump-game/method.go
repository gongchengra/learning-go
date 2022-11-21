package main

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if dp[i] == false {
			continue
		}
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i+j] = true
		}
	}
	return dp[len(nums)-1]
}
