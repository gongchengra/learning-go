package main

import (
	"container/list"
	"fmt"
	"os"
	"sort"
)

type kp struct {
	k int
	p []byte
}

func solveSudoku(board [][]byte) {
	pos := calculatePossibility(board)
	update(pos)
	if isValid(pos) {
		set(pos, board)
		return
	}
	//     printPos(pos)
	know := deepcopy(pos)
	cnt := 0
	for _, unknow := range unknowCells(pos) {
		for _, uc := range unknow.p {
			pos = deepcopy(know)
			pos[unknow.k] = []byte{uc}
			//             fmt.Println("assume ", unknow.k, string(unknow.p))
			update(pos)
			stack := list.New()
			for {
				s := status(pos)
				if s == "solved" {
					if isValid(pos) {
						if len(os.Args) > 1 {
							cnt++
							if cnt > 10 {
								printPos(pos)
								set(pos, board)
								break
							}
						}
						set(pos, board)
						return
					}
					if stack.Len() == 0 {
						break
					} else {
						v := stack.Remove(stack.Back()).([]byte)
						k := stack.Remove(stack.Back()).(int)
						pos = stack.Remove(stack.Back()).([][]byte)
						pos[k] = []byte{v[0]}
						update(pos)
						remain := remove(v, v[0])
						if len(remain) > 0 {
							stack.PushBack(deepcopy(pos))
							stack.PushBack(k)
							stack.PushBack(remain)
						}
					}
				}
				if s == "unsolved" {
					stack.PushBack(deepcopy(pos))
					k, v := leastUnknow(pos)
					remain := remove(v, v[0])
					//                     fmt.Println(k, string(v))
					stack.PushBack(k)
					stack.PushBack(remain)
					pos[k] = []byte{v[0]}
					update(pos)
				}
			}
		}
	}
}

func unknowCells(pos [][]byte) (res []kp) {
	for i, k := range pos {
		if len(k) > 1 {
			res = append(res, kp{i, k})
		}
	}
	sort.Slice(res, func(p, q int) bool {
		return len(res[p].p) > len(res[q].p)
	})
	return
}

func deepcopy(pos [][]byte) (res [][]byte) {
	for i, k := range pos {
		res = append(res, []byte{})
		for _, v := range k {
			res[i] = append(res[i], v)
		}
	}
	return
}

func printPos(pos [][]byte) {
	for i := 0; i < 81; i++ {
		if i%9 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s ", string(pos[i]))
	}
	fmt.Println()
}

func set(pos [][]byte, board [][]byte) {
	for i := 0; i < 81; i++ {
		if len(pos[i]) == 1 {
			board[i/9][i%9] = pos[i][0]
		}
	}
}

func leastUnknow(pos [][]byte) (res int, val []byte) {
	max := 10
	for i := 0; i < 81; i++ {
		//     for i := 80; i >= 0; i-- {
		if len(pos[i]) > 1 && len(pos[i]) < max {
			max, res = len(pos[i]), i
			//             return res, pos[res]
		}
	}
	// return least Unknow may lead to loop result and program will never top
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

func inPos(b byte, posi []byte) bool {
	for _, v := range posi {
		if v == b {
			return true
		}
	}
	return false
}

func update(pos [][]byte) {
	for {
		change := 0
		for i := 0; i < 81; i++ {
			if len(pos[i]) == 1 {
				for _, j := range peers(i) {
					if inPos(pos[i][0], pos[j]) && len(pos[j]) > 1 {
						pos[j] = remove(pos[j], pos[i][0])
						change++
					}
				}
			}
		}
		for i := 0; i < 81; i++ {
			if len(pos[i]) > 1 {
				for _, v := range pos[i] {
					cntc, cntr, cntb := 0, 0, 0
					for _, c := range col(i) {
						if inPos(v, pos[c]) {
							cntc++
						}
					}
					for _, r := range row(i) {
						if inPos(v, pos[r]) {
							cntr++
						}
					}
					for _, b := range block(i) {
						if inPos(v, pos[b]) {
							cntb++
						}
					}
					if cntc == 0 || cntr == 0 || cntb == 0 {
						pos[i] = []byte{v}
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
	for _, v := range pos {
		if 1 == len(v) {
			c++
		}
	}
	if 81 == c {
		return "solved"
	}
	return "unsolved"
}

func isValid(pos [][]byte) bool {
	for i := 0; i < 81; i++ {
		if len(pos[i]) != 1 {
			return false
		}
		for _, j := range peers(i) {
			if len(pos[j]) != 1 || pos[j][0] == pos[i][0] {
				return false
			}
		}
	}
	return true
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
