package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	/*
		k := 80
		fmt.Println(k, k/9, k%9)
		r, c := k/9, k%9
		r3, c3 := (r/3)*3, (c/3)*3
		fmt.Println(r, c, r3, c3)
		for j := 0; j < 9; j++ {
			cub := r3*9 + c3 + (j/3)*9 + j%3
			fmt.Println(cub)
		}
	*/
	filename := "top95.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//         fmt.Println(line)
		nums := strings.Split(line, "")
		//         fmt.Println(nums)
		b := [][]byte{}
		k := 0
		for i := 0; i < 9; i++ {
			b = append(b, []byte{})
			for j := 0; j < 9; j++ {
				b[i] = append(b[i], []byte(nums[k])[0])
				k++
			}
		}
		fmt.Println(line)
		printBoard(b)
		solveSudoku(b)
		printBoard(b)
	}
}

func printBoard(b [][]byte) {
	for _, v := range b {
		fmt.Println(string(v))
	}
}

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
