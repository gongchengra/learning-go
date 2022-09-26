package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	l := 10
	b := 10
	// Closure functions are a special case of a anonymous function
	// where you access outside variables
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
	nextInt := intSeq()
	fmt.Println(nextInt())
	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	for i := 0; i < 4; i++ {
		fmt.Println(nextInt())
		fmt.Println(nextInt2())
	}
}
