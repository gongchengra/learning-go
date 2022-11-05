package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b1, b5, b9 := randomBlk(), randomBlk(), randomBlk()
	be := []byte(".........")
	b := [][]byte{be, be, be, be, be, be, be, be, be}
	b = setValue(b, b1, 1)
	b = setValue(b, b5, 5)
	b = setValue(b, b9, 9)
	solveSudoku(b)
	removeValue(b, 60)
	printBoard(b)
}

func removeValue(board [][]byte, cnt int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cnt; {
		remove := rand.Intn(81)
		r, c := remove/9, remove%9
		if board[r][c] != '.' {
			board[r][c] = '.'
			i++
		}
	}
}

func setValue(board [][]byte, blk []byte, idx int) (res [][]byte) {
	bi := blockIdx(idx)
	for r := range board {
		res = append(res, []byte{})
		for c := range board[r] {
			res[r] = append(res[r], board[r][c])
			i := r*9 + c
			for j, b := range bi {
				if b == i {
					res[r][c] = blk[j]
				}
			}
		}
	}
	return res
}

func blockIdx(idx int) (res []int) {
	r, c := (idx-1)/3, (idx-1)%3
	for i := 0; i < 9; i++ {
		bi := r*3*9 + c*3 + (i/3)*9 + i%3
		res = append(res, bi)
	}
	return
}

func randomBlk() []byte {
	ori := []byte("123456789")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(9, func(i, j int) { ori[i], ori[j] = ori[j], ori[i] })
	return ori
}

func printBoard(b [][]byte) {
	for i, v := range b {
		fmt.Println(i, string(v))
	}
}
