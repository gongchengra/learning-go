package main

func search(nums []int, target int) bool {
	ln := len(nums)
	if ln == 0 {
		return false
	}
	k := 1
	for k < ln && nums[k-1] <= nums[k] {
		k++
	}
	i, j := 0, ln-1
	for i <= j {
		m := (i + j) / 2
		mid := (m + k) % ln
		switch {
		case nums[mid] < target:
			i = m + 1
		case nums[mid] > target:
			j = m - 1
		default:
			return true
		}
	}
	return false
}
