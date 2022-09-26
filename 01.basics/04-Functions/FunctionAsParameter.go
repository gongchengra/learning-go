package main

import "fmt"

func inc(x int) int {
	x++
	return x
}

func dec(x int) int {
	x--
	return x
}

func apply(x int, f func(int) int) int {
	r := f(x)
	return r
}

func main() {
	fmt.Println(apply(3, inc))
	fmt.Println(apply(4, dec))
}
