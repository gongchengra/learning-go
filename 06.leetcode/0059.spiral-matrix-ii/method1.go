package main

func generateMatrix(n int) [][]int {
	total, idx, res := n*n, 1, [][]int{}
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		res = append(res, tmp)
	}
	for i := 0; idx <= total; i++ {
		for j := i; idx <= total && j < n-i; j++ {
			res[i][j] = idx
			idx++
		}
		for j := i + 1; idx <= total && j < n-i; j++ {
			res[j][n-i-1] = idx
			idx++
		}
		for j := n - i - 2; idx <= total && j >= i; j-- {
			res[n-i-1][j] = idx
			idx++
		}
		for j := n - i - 2; idx <= total && j >= i+1; j-- {
			res[j][i] = idx
			idx++
		}
	}
	return res
}
