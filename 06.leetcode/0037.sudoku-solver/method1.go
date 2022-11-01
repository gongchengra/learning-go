package main

func solveSudoku(board [][]byte) {
	fill(board, '1', 0)
}

func fill(board [][]byte, n byte, block int) bool {
	if block == 9 {
		return true
	}
	if n == '9'+1 {
		return fill(board, '1', block+1)
	}
	row := (block / 3) * 3
	col := (block % 3) * 3
	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			if board[r][c] == n {
				return fill(board, n+1, block)
			}
		}
	}
	isAvailable := func(r, c int) bool {
		if board[r][c] != '.' {
			return false
		}
		for i := 0; i < 9; i++ {
			if board[r][i] == n || board[i][c] == n {
				return false
			}
		}
		return true
	}
	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			if isAvailable(r, c) {
				board[r][c] = n
				if fill(board, n+1, block) {
					return true
				}
				board[r][c] = '.'
			}
		}
	}
	return false
}
