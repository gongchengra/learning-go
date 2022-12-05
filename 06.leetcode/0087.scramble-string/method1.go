package main

/*
dp[i][j][k] 代表了s1从i开始，s2从j开始，长度为k的两个substring是否为scramble string。
有三种情况需要考虑：
如果两个substring相等的话，则为true
如果两个substring中间某一个点，左边的substrings为scramble string，同时右边的substrings也为scramble string，则为true
如果两个substring中间某一个点，s1左边的substring和s2右边的substring为scramble string, 同时s1右边substring和s2左边的substring也为scramble string，则为true
*/

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	n := len(s1)
	dp := make([][][]bool, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([][]bool, n)
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = make([]bool, n+1)
			maxK := n - max(i, j)
			for k := 1; k <= maxK; k++ {
				if s1[i:i+k] == s2[j:j+k] {
					dp[i][j][k] = true
				} else {
					for d := 1; d < k; d++ {
						if (dp[i][j][d] && dp[i+d][j+d][k-d]) ||
							(dp[i][j+k-d][d] && dp[i+d][j][k-d]) {
							dp[i][j][k] = true
							break
						}
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
