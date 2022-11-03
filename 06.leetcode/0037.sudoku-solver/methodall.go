package main

import (
	"container/list"
	"fmt"
	"sort"
)

func solveSudoku(board [][]byte) {
	pos := calculatePossibility(board)
	//     printBoard(pos)
	fmt.Println(len(peers(80)), peers(80))
	fmt.Println(len(peers(0)), peers(0))
	cnt := 0
	stack := list.New()
	for {
		s := update(pos)
		if s == "solved" {
			set(pos, board)
			cnt++
			if cnt > 9 {
				break
			}
		} else if s == "unsolved" {
			fmt.Println("unsolved")
			printPos(pos)
			stack.PushBack(pos)
			k, v := leastUnknow(pos)
			remain := remove(v, v[0])
			stack.PushBack(k)
			stack.PushBack(remain)
			fmt.Println("unsolved", k, remain)
			pos[k] = []byte{v[0]}
			update(pos)
		} else {
			fmt.Println("invalid")
			printPos(pos)
			if stack.Len() == 0 {
				break
			} else {
				v := stack.Back().Value.([]byte)
				stack.Remove(stack.Back())
				k := stack.Back().Value.(int)
				stack.Remove(stack.Back())
				pos := stack.Back().Value.([][]byte)
				stack.Remove(stack.Back())
				assume := v[0]
				pos[k] = []byte{assume}
				update(pos)
				remain := remove(v, assume)
				if len(remain) > 0 {
					stack.PushBack(k)
					stack.PushBack(remain)
					fmt.Println("invalid", k, remain)
				}
			}
		}

	}
	// printBoard(pos)
}

func printPos(pos [][]byte) {
	for i := 0; i < 81; i++ {
		if i%9 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s ", string(pos[i]))
	}
}
func set(pos [][]byte, board [][]byte) {
	for i := 0; i < 81; i++ {
		if len(pos[i]) == 1 {
			board[i/9][i%9] = pos[i][0]
		}
	}
}

func leastUnknow(pos [][]byte) (res int, val []byte) {
	max := 9
	for i := 0; i < 81; i++ {
		if len(pos[i]) > 1 && len(pos[i]) < max {
			max, res = len(pos[i]), i
		}
	}
	return res, pos[res]
}

func row(idx int) (res []int) {
	r := idx / 9
	for j := 0; j < 9; j++ {
		row := r*9 + j
		if row != idx {
			res = append(res, row)
		}
	}
	return
}

func col(idx int) (res []int) {
	c := idx % 9
	for j := 0; j < 9; j++ {
		col := j*9 + c
		if col != idx {
			res = append(res, col)
		}
	}
	return
}

func block(idx int) (res []int) {
	r, c := idx/9, idx%9
	for j := 0; j < 9; j++ {
		block := (r/3)*3*9 + (c/3)*3 + (j/3)*9 + j%3
		if block != idx {
			res = append(res, block)
		}
	}
	return
}

func peers(idx int) (res []int) {
	inSlice := func(i int) bool {
		for _, v := range res {
			if v == i {
				return true
			}
		}
		return false
	}
	for _, i := range row(idx) {
		if !inSlice(i) {
			res = append(res, i)
		}
	}
	for _, j := range col(idx) {
		if !inSlice(j) {
			res = append(res, j)
		}
	}
	for _, k := range block(idx) {
		if !inSlice(k) {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return
}

func update(pos [][]byte) (s string) {
	for {
		change := 0
		for i := 0; i < 81; i++ {
			if len(pos[i]) == 1 {
				for _, j := range peers(i) {
					inRes := func(b byte, pos []byte) bool {
						for _, v := range pos {
							if v == b {
								return true
							}
						}
						return false
					}
					if inRes(pos[i][0], pos[j]) {
						pos[j] = remove(pos[j], pos[i][0])
						change++
					}
				}
			}
		}
		if change == 0 {
			break
		}
	}
	return status(pos)
}

func remove(s []byte, c byte) (r []byte) {
	for _, v := range s {
		if v != c {
			r = append(r, v)
		}
	}
	return r
}

func status(pos [][]byte) (s string) {
	c := 0
	for i := 0; i < 81; i++ {
		if len(pos[i]) < 1 {
			return "invalid"
		} else if len(pos[i]) == 1 {
			for _, j := range peers(i) {
				if len(pos[j]) == 1 && pos[j][0] == pos[i][0] {
					return "invalid"
				}
			}
			c++
		} else {
			return "unsolved"
		}
	}
	if c == 81 {
		return "solved"
	}
	return
}

func calculatePossibility(board [][]byte) (possibility [][]byte) {
	for i := 0; i < 81; i++ {
		r := i / 9
		c := i % 9
		possibility = append(possibility, []byte{})
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
	return
}
