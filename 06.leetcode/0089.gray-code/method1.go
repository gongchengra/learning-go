package main

func grayCode(n int) []int {
	return recur(n, 1, []int{0})
}

func recur(n, base int, nums []int) []int {
	if n == 0 {
		return nums
	}
	ln := len(nums)
	tmp := make([]int, ln)
	for i := range nums {
		tmp[ln-i-1] = nums[i] + base
	}
	return recur(n-1, base*2, append(nums, tmp...))
}
