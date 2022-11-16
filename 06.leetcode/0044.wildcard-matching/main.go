package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "*"))
	fmt.Println(isMatch("cb", "c?"))
}
