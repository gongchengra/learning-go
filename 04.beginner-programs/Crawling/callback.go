package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Go program"
	c.OnRequest(func(r *colly.Request) {
		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println(r.Method)
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("-----------------------------")
		fmt.Println(e.Text)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("-----------------------------")
		fmt.Println(r.StatusCode)
		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})
	c.Visit("http://webcode.me")
}
