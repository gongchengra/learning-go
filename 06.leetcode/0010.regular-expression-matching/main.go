package main

import "fmt"

func main() {
	//     fmt.Println(isMatch("aa", "a."))
	//     fmt.Println(isMatch("aa", "a"))
	//     fmt.Println(isMatch("ab", ".a"))
	//     fmt.Println(isMatch("ab", ".b"))
	//     fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aaa", "a*"))
}

func isMatch(s string, p string) bool {
	lens, lenp := len(s), len(p)
	if lens == 0 && lenp == 0 {
		return true
	}
	dp := make([][]bool, lenp+1)
	for i := 0; i <= lenp; i++ {
		dp[i] = make([]bool, lens+1)
	}
	dp[0][0] = true
	for i := 2; i <= lenp; i += 2 {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}
	fmt.Println(dp)
	for i := 1; i <= lenp; i++ {
		if p[i-1] == '*' {
			for j := 1; j <= lens; j++ {
				dp[i][j] = dp[i-2][j] ||
					((p[i-2] == '.' || p[i-2] == s[j-1]) && dp[i][j-1])
				fmt.Println(i, j, dp)
			}
		} else {
			for j := 1; j <= lens; j++ {
				dp[i][j] = dp[i-1][j-1] && (p[i-1] == '.' || p[i-1] == s[j-1])
				fmt.Println(i, j, dp)
			}
		}
	}
	return dp[lenp][lens]
}
