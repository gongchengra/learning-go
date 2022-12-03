package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	oldname := ""
	if len(os.Args) > 1 {
		oldname = os.Args[1]
	} else {
		fmt.Println("Please provide the filename.")
		return
	}
	_, err := os.Open(oldname)
	if err == nil {
		newname := ""
		for _, runeValue := range oldname {
			if unicode.Is(unicode.Han, runeValue) || unicode.IsNumber(runeValue) || unicode.IsLetter(runeValue) || runeValue == '.' {
				newname += string(runeValue)
			}
		}
		fmt.Println("Renamed ", oldname, " to ", newname)
		e := os.Rename(oldname, newname)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		fmt.Println(err)
	}
}
