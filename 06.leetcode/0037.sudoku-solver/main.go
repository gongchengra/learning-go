package main

import "fmt"

func main() {
	if false {
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
		printBoard(b1)
		solveSudoku(b1)
		printBoard(b1)
	}
	fmt.Println()
	if true {
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
		printBoard(b2)
		solveSudoku(b2)
		printBoard(b2)
	}
}

func printBoard(b [][]byte) {
	for _, v := range b {
		fmt.Println(string(v))
	}
}

// method 2
type xy struct {
	x   int
	y   int
	all []byte
}

func solveSudoku(board [][]byte) {
	pos := []xy{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				all := []byte{}
				for n := byte('1'); n <= '9'; n++ {
					isAvailable := func(r, c int) bool {
						for k := 0; k < 9; k++ {
							if board[r][k] == n || board[k][c] == n ||
								board[(r/3)*3+(k/3)][(c/3)*3+(k%3)] == n {
								return false
							}
						}
						return true
					}
					if isAvailable(i, j) {
						all = append(all, n)
					}
				}
				pos = append(pos, xy{i, j, all})
			}
		}
	}
	fill(board, pos, 0)
}

func fill(board [][]byte, pos []xy, index int) bool {
	if index == len(pos) {
		return true
	}
	all := pos[index].all
	r := pos[index].x
	c := pos[index].y
	for i := 0; i < len(all); i++ {
		n := all[i]
		board[r][c] = n
		printBoard(board)
		valid := true
		for j := 0; j < 9; j++ {
			if (j != c && board[r][j] == n) ||
				(j != r && board[j][c] == n) ||
				((r/3)*3+j/3 != r && (c/3)*3+j%3 != c &&
					board[(r/3)*3+(j/3)][(c/3)*3+(j%3)] == n) {
				valid = false
			}
		}
		if valid && fill(board, pos, index+1) {
			return true
		}
		board[r][c] = '.'
	}
	return false
}

/* method 1
func help(board [][]byte, rows, cols, blocks [][]bool, idx int) bool {
	if idx == 81 {
		return true
	}
	i := int(idx / 9)
	j := idx % 9
	if board[i][j] != '.' {
		if help(board, rows, cols, blocks, idx+1) {
			return true
		}
	} else {
		for v := 0; v < 9; v++ {
			if rows[i][v] || cols[j][v] || blocks[i-i%3+j/3][v] {
				continue
			}
			board[i][j] = byte(v + 49)
			rows[i][v] = true
			cols[j][v] = true
			blocks[i-i%3+j/3][v] = true
			if true == help(board, rows, cols, blocks, idx+1) {
				return true
			}
			board[i][j] = '.'
			rows[i][v] = false
			cols[j][v] = false
			blocks[i-i%3+j/3][v] = false
		}
	}
	return false
}

func solveSudoku(board [][]byte) {
	var rows, cols, blocks [][]bool
	for i := 0; i < 9; i++ {
		_rows := []bool{false, false, false, false, false, false, false, false, false}
		_cols := []bool{false, false, false, false, false, false, false, false, false}
		_blocks := []bool{false, false, false, false, false, false, false, false, false}
		rows = append(rows, _rows)
		cols = append(cols, _cols)
		blocks = append(blocks, _blocks)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - 49
			rows[i][num] = true
			cols[j][num] = true
			blocks[i-i%3+j/3][num] = true
		}
	}
	help(board, rows, cols, blocks, 0)
}
*/
/* method 0

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
