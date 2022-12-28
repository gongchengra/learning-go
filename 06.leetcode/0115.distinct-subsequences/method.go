package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = 1
	}
	var prev int
	for j := 0; j < n; j++ {
		dp[j], prev = 0, dp[j]
		for i := j + 1; i < m+1; i++ {
			if t[j] == s[i-1] {
				dp[i], prev = dp[i-1]+prev, dp[i]
			} else {
				dp[i], prev = dp[i-1], dp[i]
			}
		}
	}
	return dp[m]
}
