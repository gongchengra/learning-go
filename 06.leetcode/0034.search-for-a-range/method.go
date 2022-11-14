package main

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	b, e := 0, len(nums)-1
	for b <= e {
		m := (b + e) / 2
		if nums[m] < target {
			b = m + 1
		} else if nums[m] > target {
			e = m - 1
		} else {
			i, j := m, m
			for i > 0 && nums[i-1] == nums[m] {
				i--
			}
			for j < len(nums)-1 && nums[j+1] == nums[m] {
				j++
			}
			res[0], res[1] = i, j
			break
		}
	}
	return res
}
