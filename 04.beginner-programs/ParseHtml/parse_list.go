package main

//modified from https://zetcode.com/golang/net-html/
import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileName := "test.html"
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bs)
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}
	var data []string
	doTraverse(doc, &data, "li")
	fmt.Println(data)
}

func doTraverse(doc *html.Node, data *[]string, tag string) {
	var traverse func(n *html.Node, tag string) *html.Node
	traverse = func(n *html.Node, tag string) *html.Node {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && c.Parent.Data == tag {
				*data = append(*data, c.Data)
			}
			res := traverse(c, tag)
			if res != nil {
				return res
			}
		}
		return nil
	}
	traverse(doc, tag)
}
