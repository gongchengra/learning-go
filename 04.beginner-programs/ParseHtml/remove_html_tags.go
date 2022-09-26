package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	filename := "fei.log"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	re := regexp.MustCompile(`<[^>]*>`)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	txt := string(data)
	err = os.WriteFile(filename, []byte(re.ReplaceAllString(txt, "")), 0664)
	if err != nil {
		log.Fatal(err)
	}
}
