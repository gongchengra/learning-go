package main

import (
	"fmt"
	"time"
)

func main() {
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
	b3 := [][]byte{
		[]byte("4.....8.5"),
		[]byte(".3......."),
		[]byte("...7....."),
		[]byte(".2.....6."),
		[]byte("....8.4.."),
		[]byte("....1...."),
		[]byte("...6.3.7."),
		[]byte("5..2....."),
		[]byte("1.4......"),
	}
	if false {
		printBoard(b1)
		//         solveSudoku(b1)
		//         printBoard(b1)
		printBoard(b2)
	}
	//         solveSudoku(b2)
	//         printBoard(b2)
	//         printBoard(b3)
	// method1 332.251244ms
	// method2 617.430751ms
	// method3 3.283633214s
	// methodall 62.988252ms
	t0 := time.Now()
	solveSudoku(b3)
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	printBoard(b3)
}

func printBoard(b [][]byte) {
	for i, v := range b {
		fmt.Println(i, string(v))
	}
}
