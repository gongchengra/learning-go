package main

import "fmt"

func main() {
	//     m := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}
	m := [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}
	setZeroes(m)
	fmt.Println(m)
}
