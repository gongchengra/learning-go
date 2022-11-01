package main

import "sort"

type xy struct {
	x   int
	y   int
	all []byte
}

func simplify(board [][]byte) {
	for {
		change := 0
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
					if len(all) == 1 {
						board[i][j] = all[0]
						change++
					}
				}
			}
		}
		if change == 0 {
			break
		}
	}
}

func updatePossibility(board [][]byte) (pos []xy) {
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
	return
}

func solveSudoku(board [][]byte) {
	//     simplify(board)
	pos := updatePossibility(board)
	sort.Slice(pos, func(p, q int) bool {
		return len(pos[p].all) < len(pos[q].all)
	})
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

func remove(s []byte, c byte) (r []byte) {
	for _, v := range s {
		if v != c {
			r = append(r, v)
		}
	}
	return r
}
