package main

import "fmt"

func main() {
	t := Ints2TreeNode([]int{1, 3, 0, 0, 2})
	recoverTree(t)
	fmt.Println(Tree2Inorder(t))
}
