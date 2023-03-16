package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 要检测字符串
	fmt.Println("请输入一个字符串")
	var str string
	fmt.Scan(&str)
	letterRegex := regexp.MustCompile("[a-zA-Z]")
	numberRegex := regexp.MustCompile("[0-9]")
	specialRegex := regexp.MustCompile("[^a-zA-Z0-9]")
	isValid := len(str) >= 8 && letterRegex.MatchString(str) && numberRegex.MatchString(str) && specialRegex.MatchString(str)
	if isValid {
		fmt.Printf("%s satisfies the condition\n", str)
	} else {
		fmt.Printf("%s does not satisfy the condition\n", str)
	}
}
