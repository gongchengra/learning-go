package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	filename := "words.log"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	envurl := os.Getenv("URL")
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
		searchUrl := envurl + word
		resp, err := http.Get(searchUrl)
		if err != nil {
			log.Fatal(err)
		}
		body, _ := io.ReadAll(resp.Body)
		if len(body) == 0 {
			fmt.Println("Lookup for ", word)
			uri := baseUrl + word
			content := curl(uri)
			if content != "" {
				data := url.Values{
					"def": {(word + ":\n" + content)},
				}
				resp, err = http.PostForm(searchUrl, data)
				if err != nil {
					log.Fatal(err)
				}
			}
			n := time.Duration(5 + rand.Intn(25))
			time.Sleep(n * time.Second)
			c++
			if c > 0 && c%10 == 0 {
				time.Sleep(n * time.Minute)
			}
		}
	}
}
