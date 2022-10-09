package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

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
	f, err := os.Open(".")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	if dirs, err := f.Readdir(-1); err == nil {
		for _, d := range dirs {
			filename := d.Name()
			match, _ := regexp.Match(`^[0-9a-zA-Z]`, []byte(filename))
			if match == false {
				re := regexp.MustCompile(`[\pP]+`)
				ext := filepath.Ext(filename)
				file := strings.TrimSuffix(filename, ext)
				sp := regexp.MustCompile(`\s`)
				clean := sp.ReplaceAllString(re.ReplaceAllString(file, ""), "")
				newname := strToRuneSumString(file) + clean + ext
				fmt.Println("Renamed ", filename, " to ", newname)
				e := os.Rename(filename, newname)
				if e != nil {
					fmt.Println(e)
				}
			}
		}
	}
}
