package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	urls := []string{
		"http://webcode.me",
		"https://example.com",
		"http://httpbin.org",
		"https://www.perl.org",
		"https://www.php.net",
		"https://www.python.org",
		"https://code.visualstudio.com",
		"https://clojure.org",
	}
	c := colly.NewCollector(
		colly.Async(),
	)
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	for _, url := range urls {
		c.Visit(url)
	}
	c.Wait()
}
