package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j > 0 {
					grid[0][j] += grid[0][j-1]
				}
				continue
			}
			if j == 0 {
				if i > 0 {
					grid[i][0] += grid[i-1][0]
				}
				continue
			}
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[m-1][n-1]
}
