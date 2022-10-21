package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	urls := []string{}
	filename := "t.log"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	c := colly.NewCollector(colly.MaxDepth(1))
	c.OnHTML("pre", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	for i := len(urls) - 1; i >= 0; i-- {
		c.Visit(urls[i])
		//         fmt.Println(urls[i])
	}
}
