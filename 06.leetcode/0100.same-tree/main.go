package main

import "fmt"

func main() {
	t1 := Ints2TreeNode([]int{1, 2, 3})
	t2 := Ints2TreeNode([]int{1, 2, 3})
	fmt.Println(isSameTree(t1, t2))
	t3 := Ints2TreeNode([]int{1, 2})
	t4 := Ints2TreeNode([]int{1, 0, 2})
	fmt.Println(isSameTree(t3, t4))
}
