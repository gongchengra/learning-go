package main

import "fmt"

func main() {
	fmt.Println(isValidBST(Ints2TreeNode([]int{2, 1, 3})))
	fmt.Println(isValidBST(Ints2TreeNode([]int{5, 1, 4, 0, 0, 3, 6})))
}
