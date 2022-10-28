package main

import "fmt"

func main() {
	b := [][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	fmt.Println(isValidSudoku(b))
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
	fmt.Println(isValidSudoku(b1))
	b2 := [][]byte{
		[]byte("83..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	fmt.Println(isValidSudoku(b1))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			} else {
				for k := 0; k < 9; k++ {
					if k != i && board[k][j] == board[i][j] {
						return false
					}
				}
				for k := 0; k < 9; k++ {
					if k != j && board[i][k] == board[i][j] {
						return false
					}
				}
				ik := i / 3
				jk := j / 3
				for k := 0; k < 9; k++ {
					if !(k/3+ik*3 == i && k%3+jk*3 == j) &&
						board[k/3+ik*3][k%3+jk*3] == board[i][j] {
						return false
					}
				}
			}
		}
	}
	return true
}
