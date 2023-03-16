package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("请输入一个字符串")
	var str string
	fmt.Scan(&str)
	newString := ""
	for _, char := range str {
		if unicode.IsDigit(char) || unicode.IsLetter(char) || unicode.Is(unicode.Han, char) {
			newString += string(char)
		}
	}
	fmt.Println(strings.TrimSpace(newString))
}
