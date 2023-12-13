package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSudoku(board [][]int) bool {
	emptyCell := findEmptyCell(board)
	if emptyCell == nil {
		return true
	}
	for num := 1; num <= 9; num++ {
		if isValidPlacement(board, emptyCell, num) {
			board[emptyCell[0]][emptyCell[1]] = num
			if solveSudoku(board) {
				return true
			}
			board[emptyCell[0]][emptyCell[1]] = 0
		}
	}
	return false
}

func findEmptyCell(board [][]int) []int {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return []int{row, col}
			}
		}
	}
	return nil
}

func isValidPlacement(board [][]int, cell []int, num int) bool {
	row, col := cell[0], cell[1]
	// Check row
	for i := 0; i < 9; i++ {
		if board[row][i] == num && i != col {
			return false
		}
	}
	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num && i != row {
			return false
		}
	}
	// Check subgrid
	subgridRowStart := (row / 3) * 3
	subgridColStart := (col / 3) * 3
	for i := subgridRowStart; i < subgridRowStart+3; i++ {
		for j := subgridColStart; j < subgridColStart+3; j++ {
			if board[i][j] == num && (i != cell[0] || j != cell[1]) {
				return false
			}
		}
	}
	return true
}

func main() {
	board, err := readBoardFromFile("input.log")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	solvedBoard := copyBoard(board)
	if solveSudoku(solvedBoard) {
		fmt.Println("Solved Sudoku:")
		for _, row := range solvedBoard {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Sudoku could not be solved.")
	}
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, 9)
	for i := 0; i < 9; i++ {
		row := make([]int, 9)
		copy(row, board[i])
		newBoard[i] = row
	}
	return newBoard
}

func readBoardFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var board [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		board = append(board, row)
	}
	return board, nil
}
