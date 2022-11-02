package main

import (
	"container/list"
	"fmt"
	"sort"
)

func solveSudoku(board [][]byte) {
	pos := calculatePossibility(board)
	//     printBoard(pos)
	cnt := 0
	posStack := list.New()
	keyStack := list.New()
	valStack := list.New()
	for {
		update(pos)
		s := status(pos)
		if s == "solved" {
			printBoard(pos)
			set(pos, board)
			cnt++
			if cnt > 9 {
				break
			}
		} else if s == "unsolved" {
			posStack.PushBack(pos)
			idx := leastUnknow(pos)
			keyStack.PushBack(idx)
			assume := pos[idx][0]
			remain := remove(pos[idx], assume)
			valStack.PushBack(remain)
			fmt.Println("\nPushed ", idx, string(remain))
			fmt.Println("Unsolved Assume postion ", idx, " to be ", string(assume))
			pos[idx] = []byte{assume}
			printPos(pos)
			update(pos)
			//             set(pos, board)
			//             printBoard(board)
		} else {
			if posStack.Len() == 0 {
				break
			}
			pos = posStack.Back().Value.([][]byte)
			posStack.Remove(posStack.Back())
			if keyStack.Len() == 0 {
				break
			}
			k := keyStack.Back().Value.(int)
			keyStack.Remove(keyStack.Back())
			if valStack.Len() == 0 {
				break
			}
			v := valStack.Back().Value.([]byte)
			valStack.Remove(valStack.Back())
			assume := v[0]
			pos[k] = []byte{assume}
			remain := remove(v, assume)
			if len(remain) > 0 {
				posStack.PushBack(pos)
				valStack.PushBack(remain)
			}
			fmt.Println("\nPushed ", k, string(remain))
			fmt.Println("Invalid: Assume postion ", k, " to be ", string(assume))
			printPos(pos)
			update(pos)
			//             set(pos, board)
			//             printBoard(board)
			/*
				kv := kvStack.Back().Value.(map[int][]byte)
				fmt.Println(kv)
				for k, v := range kv {
					pos[k] = []byte{v[0]}
				}
			*/
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

func leastUnknow(pos [][]byte) (res int) {
	max := 9
	for i := 0; i < 81; i++ {
		if len(pos[i]) > 1 && len(pos[i]) < max {
			max, res = len(pos[i]), i
		}
	}
	return
}

func peers(idx int) (res []int) {
	r, c := idx/9, idx%9
	for j := 0; j < 9; j++ {
		inSlice := func(i int) bool {
			for _, v := range res {
				if v == i {
					return true
				}
			}
			return false
		}
		row := r*9 + j
		if row != idx && !inSlice(row) {
			res = append(res, row)
		}
		col := j*9 + c
		if col != idx && !inSlice(col) {
			res = append(res, col)
		}
		block := (r/3)*3*9 + (c/3)*3 + (j/3)*9 + j%3
		if block != idx && !inSlice(block) {
			res = append(res, block)
		}
	}
	sort.Ints(res)
	return
}

func update(pos [][]byte) {
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
