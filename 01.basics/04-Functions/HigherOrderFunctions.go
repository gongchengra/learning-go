package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

// Function that returns another function

func partialMultiplication(x int) func(int) int {
	return func(y int) int {
		return multiply(x, y)
	}
}

func apply(x, y int, add func(int, int) int, sub func(int, int) int) (int, int) {
	return add(x, y), sub(x, y)
}

func getAddSub() (func(int, int) int, func(int, int) int) {
	return func(x, y int) int { return x + y }, func(x, y int) int { return x - y }
}

func main() {
	multiple := partialMultiplication(10)
	fmt.Println(multiple(10))
	x := 3
	y := 4
	add, sub := getAddSub()
	r1, r2 := apply(x, y, add, sub)
	fmt.Printf("%d + %d = %d\n", x, y, r1)
	fmt.Printf("%d - %d = %d\n", x, y, r2)
}
