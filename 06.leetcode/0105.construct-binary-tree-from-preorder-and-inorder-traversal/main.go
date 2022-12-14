package main

import "fmt"

func main() {
	fmt.Println(Tree2Preorder(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
