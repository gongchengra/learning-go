package main

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	if n < 2 {
		return matrix[0]
	}
	m := len(matrix[0])
	ret := make([]int, m*n)
	idx, max := 0, m*n
	for i := 0; idx < max; i++ {
		for j := i; j < m-i && idx < max; j++ {
			ret[idx] = matrix[i][j]
			idx++
		}
		for j := i + 1; j < n-i && idx < max; j++ {
			ret[idx] = matrix[j][m-i-1]
			idx++
		}
		for j := m - i - 2; j >= i && idx < max; j-- {
			ret[idx] = matrix[n-i-1][j]
			idx++
		}
		for j := n - i - 2; j > i && idx < max; j-- {
			ret[idx] = matrix[j][i]
			idx++
		}
	}
	return ret
}
