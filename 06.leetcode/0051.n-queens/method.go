package main

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	cols := make([]bool, n)
	slash := make([]bool, 2*n)
	backslash := make([]bool, 2*n)
	board := make([]string, n)
	res := [][]string{}
	dfs(0, cols, slash, backslash, board, &res)
	return res
}

func dfs(row int, cols, slash, backslash []bool, board []string, res *[][]string) {
	if row == len(board) {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*res = append(*res, tmp)
		return
	}
	n := len(board)
	for c := 0; c < len(board); c++ {
		si := row - c + n
		bi := 2*n - row - c - 1
		if !cols[c] && !slash[si] && !backslash[bi] {
			b := make([]byte, n)
			for i := range b {
				b[i] = '.'
			}
			b[c] = 'Q'
			board[row] = string(b)
			cols[c], slash[si], backslash[bi] = true, true, true
			dfs(row+1, cols, slash, backslash, board, res)
			cols[c], slash[si], backslash[bi] = false, false, false
		}
	}
}
