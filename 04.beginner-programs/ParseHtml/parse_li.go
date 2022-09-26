package main

//modified from https://zetcode.com/golang/net-html/
import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
)

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func parse(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var vals []string
	var isLi bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals
		case tt == html.StartTagToken:
			t := tkn.Token()
			isLi = t.Data == "li"
		case tt == html.TextToken:
			t := tkn.Token()
			if isLi {
				vals = append(vals, t.Data)
			}
			isLi = false
		}
	}
}

func main() {
	fileName := "test.html"
	text, err := readHtmlFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	data := parse(text)
	fmt.Println(data)
}
