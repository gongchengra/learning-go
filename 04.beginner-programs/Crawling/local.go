package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
)

func main() {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	c := colly.NewCollector()
	c.WithTransport(t)
	words := []string{}
	c.OnHTML("li", func(e *colly.HTMLElement) {
		words = append(words, e.Text)
	})
	c.Visit("file://./words.html")
	for _, p := range words {
		fmt.Printf("%s\n", p)
	}
}
