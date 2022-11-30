package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 13))
}
