package main

func trap(height []int) int {
	res, l, r, maxl, maxr := 0, 0, len(height)-1, 0, 0
	for l <= r {
		if height[l] <= height[r] {
			if height[l] > maxl {
				maxl = height[l]
			} else {
				res += maxl - height[l]
			}
			l++
		} else {
			if height[r] > maxr {
				maxr = height[r]
			} else {
				res += maxr - height[r]
			}
			r--
		}
	}
	return res
}
