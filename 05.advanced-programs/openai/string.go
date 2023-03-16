package main

import (
	"fmt"
	"strings"
)

func main() {
	//输入的字符串
	str := "abcdefghij"
	//处理一：将字符串逆序
	revStr := Reverse(str)
	//处理二：将第3,6,9位字母转换成大写
	upperStr := Exchange(revStr)
	//处理三：截取字符串的前7位
	finStr := upperStr[:7]
	//输出
	fmt.Println(finStr)
}

// Reverse 将字符串逆序

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Exchange 第3,6,9位字母转换成大写

func Exchange(s string) string {
	strs := strings.Split(s, "")
	strs[2] = strings.ToUpper(strs[2])
	strs[5] = strings.ToUpper(strs[5])
	strs[8] = strings.ToUpper(strs[8])
	return strings.Join(strs, "")
}
