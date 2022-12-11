package main

// solution by chatgpt
// Interleaved returns whether s3 is formed by an interleaving of s1 and s2.
func isInterleave(s1, s2, s3 string) bool {
	// If the length of s3 is not equal to the sum of the lengths of s1 and s2,
	// it is not possible for s3 to be formed by an interleaving of s1 and s2.
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	// Create a 2D array of booleans to store the intermediate results of our
	// dynamic programming algorithm. The array will have a size of (len(s1)+1)
	// by (len(s2)+1) because we need to account for the empty string as well.
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// Set the initial values for the array.
	// The empty string is always an interleaving of the empty string, so dp[0][0]
	// is set to true.
	// If s1 is empty, dp[0][j] is true if and only if the first j characters of
	// s2 match the first j characters of s3.
	// If s2 is empty, dp[i][0] is true if and only if the first i characters of
	// s1 match the first i characters of s3.
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j <= len(s2); j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Use dynamic programming to fill in the rest of the array.
	// dp[i][j] is true if and only if either:
	// - the first i characters of s1 match the first i+j characters of s3 and the
	//   (i+1)th character of s3 is equal to the (j+1)th character of s2, or
	// - the first j characters of s2 match the first i+j characters of s3 and the
	//   (j+1)th character of s3 is equal to the (i+1)th character of s1.
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] && s2[j-1] != s3[i+j-1] {
				dp[i][j] = dp[i-1][j]
			} else if s2[j-1] == s3[i+j-1] && s1[i-1] != s3[i+j-1] {
				dp[i][j] = dp[i][j-1]
			} else if s1[i-1] == s3[i+j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	// Return the final value of dp[len(s1)][len(s2)].
	return dp[len(s1)][len(s2)]
}
