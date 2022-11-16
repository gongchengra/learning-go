package main

func jump(nums []int) int {
	i, cnt, end := 0, 0, len(nums)-1
	next, maxNext, maxI := 0, 0, 0
	for i < end {
		if i+nums[i] >= end {
			return cnt + 1
		}
		next, maxNext = i+1, i+nums[i]
		for next <= maxNext {
			if next+nums[next] > maxI {
				maxI, i = next+nums[next], next
			}
			next++
		}
		cnt++
	}
	return cnt
}
