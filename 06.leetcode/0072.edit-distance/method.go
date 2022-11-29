package main

func minDistance(word1 string, word2 string) int {
	m, n, cost := len(word1), len(word2), 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			if word1[i-1] == word2[j-1] {
				cost = 0
			} else {
				cost = 1
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+cost)
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
