package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m == 0 || n == 0 || len(word) == 0 {
		return false
	}
	var dfs func(int, int, int) bool
	dfs = func(r, c, idx int) bool {
		if len(word) == idx {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[idx] {
			return false
		}
		tmp := board[r][c]
		board[r][c] = 0
		if dfs(r-1, c, idx+1) || dfs(r+1, c, idx+1) || dfs(r, c-1, idx+1) || dfs(r, c+1, idx+1) {
			return true
		}
		board[r][c] = tmp
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
