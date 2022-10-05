package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type question struct {
	title   string
	excerpt string
}

func main() {
	c := colly.NewCollector()
	qs := []question{}
	c.OnHTML("div.s-post-summary", func(e *colly.HTMLElement) {
		q := question{}
		q.title = e.ChildText("a.s-link")
		q.excerpt = e.ChildText(".s-post-summary--content-excerpt")
		qs = append(qs, q)
	})
	c.OnScraped(func(r *colly.Response) {
		for idx, q := range qs {
			fmt.Println("---------------------------------")
			fmt.Println(idx + 1)
			fmt.Printf("Q: %s \n\n", q.title)
			fmt.Println(q.excerpt)
		}
	})
	c.Visit("https://stackoverflow.com/questions/tagged/perl")
}
