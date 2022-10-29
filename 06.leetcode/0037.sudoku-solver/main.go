package main

import "fmt"

func main() {
	if true {
		b1 := [][]byte{
			[]byte("53..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		}
		solveSudoku(b1)
		printBoard(b1)
	}
	if false {
		b2 := [][]byte{
			[]byte("........."),
			[]byte(".23...78."),
			[]byte("1..4.6..9"),
			[]byte("4...5...1"),
			[]byte("9.......6"),
			[]byte(".6.....9."),
			[]byte("..5...8.."),
			[]byte("...3.1..."),
			[]byte("....9...."),
		}
		solveSudoku(b2)
		printBoard(b2)
	}
}

func printBoard(b [][]byte) {
	for _, v := range b {
		fmt.Println(string(v))
	}
}

func updatePossibility(board [][]byte, possibility [][]byte) {
	for i := 0; i < 81; i++ {
		r := i / 9
		c := i % 9
		possibility[i] = []byte{}
		if board[r][c] != '.' {
			possibility[i] = append(possibility[i], board[r][c])
		} else {
			for n := byte('1'); n <= '9'; n++ {
				isAvailable := func(r, c int) bool {
					for i := 0; i < 9; i++ {
						if board[r][i] == n || board[i][c] == n ||
							board[(r/3)*3+(i/3)][(c/3)*3+(i%3)] == n {
							return false
						}
					}
					return true
				}
				if isAvailable(r, c) {
					possibility[i] = append(possibility[i], n)
				}
			}
		}
	}
}

func solveSudoku(board [][]byte) {
	possibility := [][]byte{}
	for i := 0; i < 81; i++ {
		possibility = append(possibility, []byte{})
	}
	updatePossibility(board, possibility)
	for {
		change := 0
		for i := 0; i < 81; i++ {
			r := i / 9
			c := i % 9
			if board[r][c] == '.' && len(possibility[i]) == 1 {
				board[r][c] = possibility[i][0]
				change++
				updatePossibility(board, possibility)
			}
		}
		if change == 0 {
			break
		}
	}
}

func hasSolved(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

func remove(s []byte, c byte) (r []byte) {
	for _, v := range s {
		if v != c {
			r = append(r, v)
		}
	}
	return r
}

/*
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
*/
