package main

func climbStairs(n int) int {
	s := make([]int, n+1)
	s[0], s[1] = 1, 1
	for i := 2; i < n+1; i++ {
		s[i] = s[i-2] + s[i-1]
	}
	return s[n]
}
