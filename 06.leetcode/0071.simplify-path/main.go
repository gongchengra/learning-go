package main

import "fmt"

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo//bar"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
