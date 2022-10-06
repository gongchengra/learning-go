package main

import (
	"bufio"
	"database/sql"
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
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		log.Fatal(err)
	}
	wordDb := Newsqlitedb(db)
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
		_, err := wordDb.GetByWord(word)
		if err != nil {
			if err == ErrNotExists {
				fmt.Println("Lookup for ", word)
				uri := baseUrl + word
				content := curl(uri)
				if content != "" {
					_, err = wordDb.Create(word, (word + ":\n" + content))
				}
				n := time.Duration(5 + rand.Intn(25))
				time.Sleep(n * time.Second)
				c++
				if c > 0 && c%100 == 0 {
					time.Sleep(n * time.Minute)
				}
			}
		}
	}
}
