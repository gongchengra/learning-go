package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.Visit("http://webcode.me")
}
