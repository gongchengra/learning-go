package main

import "fmt"

func main() {
	tree := Ints2TreeNode([]int{1, 2, 5, 3, 4, NULL, 6})
	flatten(tree)
	fmt.Println(Tree2Inorder(tree))
}
