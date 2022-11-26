package main

import "fmt"

func main() {
	fmt.Println(isNumber("0"), isNumber("."), isNumber("e"))
	fmt.Println(isNumber("0089"), isNumber("-0.1"), isNumber("-80e30"))
}
