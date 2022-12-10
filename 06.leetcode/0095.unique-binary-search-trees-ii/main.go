package main

import "fmt"

func main() {
	for _, v := range generateTrees(3) {
		fmt.Println(Tree2Preorder(v))
	}
}
