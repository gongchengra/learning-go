package main

func getPermutation(n int, k int) string {
	size := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	res := ""
	k--
	for i := 0; i < n; i++ {
		total := size[n-i]
		idx := int(k / int(total/(n-i)))
		idx %= n - i
		t := nums[idx]
		nums = append(nums[:idx], nums[idx+1:]...)
		res += string('0' + t)
	}
	return res
}
