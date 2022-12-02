package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	oldname := "hello, 世界"
	if len(os.Args) > 1 {
		oldname = os.Args[1]
	}
	newname := ""
	for _, runeValue := range oldname {
		if unicode.Is(unicode.Han, runeValue) || unicode.IsLetter(runeValue) || unicode.IsLetter(runeValue) || runeValue == '.' {
			newname += string(runeValue)
		}
	}
	fmt.Println(newname)
}
