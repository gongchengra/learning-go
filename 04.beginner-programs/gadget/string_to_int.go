package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func strToInt(s string) int {
	r := 0
	for _, i := range []byte(s) {
		r += int(i)
	}
	return r
}

func strToIntString(s string) string {
	r := ""
	for _, i := range []byte(s) {
		r += strconv.Itoa(int(i))
	}
	return r
}

func strToRuneString(s string) string {
	res := ""
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		res += strconv.Itoa(int(r))
		s = s[size:]
	}
	return res
}

func strToRuneSumString(s string) string {
	res := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		res += int(r)
		s = s[size:]
	}
	return strconv.Itoa(int(res))
}

func main() {
	s := "八仙得道传"
	fmt.Println(s)
	//     fmt.Println(utf8.DecodeRuneInString(s))
	//     fmt.Println([]byte(s))
	fmt.Println(strToInt(s))
	fmt.Println(strToIntString(s))
	fmt.Println(strToRuneString(s))
	fmt.Println(strToRuneSumString(s))
	fmt.Println(utf8.RuneCountInString(s))
	{
		q := "清"
		r, _ := utf8.DecodeRuneInString(q)
		fmt.Printf("%b\n", r)
	}
	{
		q := []byte("清")
		r, _ := utf8.DecodeRune(q)
		fmt.Printf("%b %[1]d %[1]x\n", r)
	}
}
