package main

import "fmt"

func main() {
	fmt.Println(maxPathSum(Ints2TreeNode([]int{1, 2, 3})))
	fmt.Println(maxPathSum(Ints2TreeNode([]int{-10, 9, 20, NULL, NULL, 15, 7})))
}
