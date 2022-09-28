package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func lessPrint(c string) {
	r := strings.Split(c, "\n")
	var input int
	for i := range r {
		if i > 0 && i%20 == 0 {
			fmt.Scanf("%d", &input)
		}
		fmt.Println(i, r[i])
	}
}

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	word := "go"
	if len(os.Args) > 1 {
		word = os.Args[1]
	}
	filename := word + ".txt"
	if fileExists(filename) {
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		lessPrint(string(content))
	} else {
		uri := baseUrl + word
		content := curl(uri)
		if content != "" {
			f, err := os.Create(word + ".txt")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			f.WriteString(word + ":\n")
			f.WriteString(content)
			lessPrint(content)
		}
	}
}
