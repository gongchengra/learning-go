package main

// From https://zetcode.com/golang/net-html/
import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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
	showTitles(urls)
}

func showTitles(urls []string) {
	c := getTitleTags(urls)
	for msg := range c {
		fmt.Println(msg)
	}
}

func getTitleTags(urls []string) chan string {
	c := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go getTitle(url, c)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func getTitle(url string, c chan string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		c <- "failed to fetch data"
		return
	}
	defer resp.Body.Close()
	tkn := html.NewTokenizer(resp.Body)
	var isTitle bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			isTitle = t.Data == "title"
		case tt == html.TextToken:
			t := tkn.Token()
			if isTitle {
				c <- t.Data
				isTitle = false
			}
		}
	}
}
