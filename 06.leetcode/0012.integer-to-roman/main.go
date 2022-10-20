package main

import "fmt"

func main() {
	//     fmt.Println(intToRoman(55))
	for i := 1; i < 4000; i++ {
		fmt.Println(i, intToRoman(i))
	}
}

func intToRoman(num int) string {
	symbol := [...]string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := ""
	for i := 0; num != 0; i++ {
		for num >= value[i] {
			num -= value[i]
			str += symbol[i]
		}
	}
	return str
}
