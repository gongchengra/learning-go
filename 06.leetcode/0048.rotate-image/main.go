package main

import "fmt"

func main() {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(m)
	rotate(m)
	fmt.Println(m)
	m1 := [][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}}
	fmt.Println(m1)
	rotate(m1)
	fmt.Println(m1)
}
