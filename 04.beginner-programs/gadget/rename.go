package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func strToRuneSumString(s string) string {
	res := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		res += int(r)
		s = s[size:]
	}
	return strconv.Itoa(res)
}

func cleanupName(filename string) string {
	newname := ""
	for _, runeValue := range filename {
		if unicode.Is(unicode.Han, runeValue) || unicode.IsNumber(runeValue) || unicode.IsLetter(runeValue) || runeValue == '.' || runeValue == '_' || runeValue == '-' {
			newname += string(runeValue)
		}
	}
	return newname
}

func renameFile(filename string) error {
	newname := cleanupName(filename)
	ext := filepath.Ext(newname)
	file := strings.TrimSuffix(newname, ext)
	match, _ := regexp.Match(`^[0-9a-zA-Z]`, []byte(filename))
	if match == false {
		newname = strToRuneSumString(file) + file + ext
	}
	if filename != newname {
		fmt.Println("Renaming", filename, "to", newname)
		err := os.Rename(filename, newname)
		return err
	}
	return nil
}

func main() {
	if len(os.Args) == 2 {
		// Single file operation mode
		filename := os.Args[1]
		err := renameFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Batch rename operation mode
		f, err := os.Open(".")
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		dirs, err := f.Readdir(-1)
		if err != nil {
			log.Println(err)
			return
		}

		for _, d := range dirs {
			if d.IsDir() {
				continue // skip directories
			}
			filename := d.Name()
			err := renameFile(filename)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
