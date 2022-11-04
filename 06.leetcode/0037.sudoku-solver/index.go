package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i, v := range b {
		fmt.Println(i, string(v))
	}
}
