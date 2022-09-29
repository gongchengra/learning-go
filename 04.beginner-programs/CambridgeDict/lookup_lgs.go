package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	search := "go"
	if len(os.Args) > 1 {
		search = os.Args[1]
	}
	envurl := os.Getenv("URL")
	searchUrl := envurl + search
	resp, err := http.Get(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		uri := baseUrl + search
		content := curl(uri)
		if content != "" {
			data := url.Values{
				"def": {(search + ":\n" + content)},
			}
			resp, err = http.PostForm(searchUrl, data)
			if err != nil {
				log.Fatal(err)
			}
			lessPrint(content)
		}
	} else {
		lessPrint(string(body))
	}
}
