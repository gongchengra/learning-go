package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		res := ""
		char := str[0]
		count := 1
		for i := 1; i < len(str); i++ {
			if char != str[i] {
				res += string(byte(count + '0'))
				res += string(char)
				count = 1
				char = str[i]
			} else {
				count++
			}
		}
		res += string(byte(count + '0'))
		res += string(char)
		str = res
	}
	return str
}
