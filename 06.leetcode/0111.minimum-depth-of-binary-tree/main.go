package main

import "fmt"

func main() {
	fmt.Println(minDepth(Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7})))
	fmt.Println(minDepth(Ints2TreeNode([]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6})))
}
