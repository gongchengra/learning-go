package main

import (
	"fmt"
	"os"
	"strconv"
)

// Function to divide numbers represented as arrays by a given divisor
func divideArray(array []int, divisor int) {
	modulo := 0
	for i := 0; i < len(array); i++ {
		tmp := array[i] + modulo*10
		array[i] = tmp / divisor
		modulo = tmp % divisor
	}
}

// Print the array as a number, formatting it to print 64 digits per line
func printArray(array []int) {
	start := len(array) - 1
	for start >= 0 && array[start] == 0 {
		start--
	}
	if start == -1 {
		fmt.Print("0")
	} else {
		count := 0
		for i := 1; i <= start; i++ {
			fmt.Print(array[i])
			count++
			if count%64 == 0 {
				fmt.Println()
			}
		}
		if count%64 != 0 {
			fmt.Println()
		}
	}
}

func copyArray(source, target []int) {
	for i := range source {
		target[i] = source[i]
	}
}

// Add two arrays and store the result in a third array
func addArrays(augend, addend, sum []int) {
	carry := 0
	for i := len(sum) - 1; i >= 0; i-- {
		sum[i] = augend[i] + addend[i] + carry
		if sum[i] >= 10 {
			sum[i] -= 10
			carry = 1
		} else {
			carry = 0
		}
	}
}

// Subtract one array from another
func subtractArrays(minuend, subtractor, result []int) {
	borrow := 0
	for i := len(result) - 1; i >= 0; i-- {
		if minuend[i]-borrow >= subtractor[i] {
			result[i] = minuend[i] - subtractor[i] - borrow
			borrow = 0
		} else {
			result[i] = 10 + minuend[i] - subtractor[i] - borrow
			borrow = 1
		}
	}
}

func calculatePi(length int) {
	// Create slices based on the length provided
	pi := make([]int, length)
	d5 := make([]int, length)
	d239 := make([]int, length)
	t5 := make([]int, length)
	t239 := make([]int, length)

	d5[0] = 16
	d239[0] = 4
	divideArray(d5, 5)
	divideArray(d239, 239)

	flag := 1
	for i := 1; i < length*3/2; i += 2 {
		copyArray(d5, t5)
		copyArray(d239, t239)
		divideArray(t5, i)
		divideArray(t239, i)

		if flag > 0 {
			addArrays(pi, t5, pi)
			subtractArrays(pi, t239, pi)
		} else {
			subtractArrays(pi, t5, pi)
			addArrays(pi, t239, pi)
		}
		flag = -flag
		divideArray(d5, 25)
		divideArray(d239, 239*239)
	}

	printArray(pi)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the desired length.")
		return
	}
	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length <= 0 {
		fmt.Println("Please provide a valid positive integer.")
		return
	}
	calculatePi(length)
}
