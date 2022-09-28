package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	filename := "words.log"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	rand.Seed(time.Now().UnixNano())
	c := 0
	for fileScanner.Scan() {
		word := fileScanner.Text()
		if !fileExists(word + ".txt") {
			fmt.Println("Lookup for ", word)
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
				n := time.Duration(1 + rand.Intn(5))
				time.Sleep(n * time.Second)
				c++
				if c > 0 && c%100 == 0 {
					time.Sleep(n * time.Minute)
				}
			}
		}
	}
}
