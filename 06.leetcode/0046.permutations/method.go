package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i := range nums {
		s := []int{}
		s = append(s, nums[0:i]...)
		s = append(s, nums[i+1:]...)
		for _, v := range permute(s) {
			tmp := []int{}
			tmp = append(tmp, nums[i])
			tmp = append(tmp, v...)
			res = append(res, tmp)
		}
	}
	return res
}
