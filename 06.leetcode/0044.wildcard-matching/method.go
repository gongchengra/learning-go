package main

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := [][]bool{}
	for i := 0; i < ls+1; i++ {
		rt := make([]bool, lp+1)
		dp = append(dp, rt)
	}
	dp[0][0] = true
	for j := 1; j <= lp; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '?') && dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[ls][lp]
}
