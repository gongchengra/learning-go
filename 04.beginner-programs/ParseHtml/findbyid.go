package main

// go build findbyid.go util.go
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return ""
	}
	return buf.String()
}

func checkId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string) *html.Node {
	if checkId(n, id) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := traverse(c, id)
		if res != nil {
			return res
		}
	}
	return nil
}

func getElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id)
}

func main() {
	f, err := os.Open(".")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	var vals []string
	if dirs, err := f.Readdir(-1); err == nil {
		SortFileNameAscend(dirs)
		for _, d := range dirs {
			filename := d.Name()
			if !strings.Contains(filename, "html") || strings.Contains(filename, "index") {
				continue
			}
			text, err := os.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			doc, err := html.Parse(strings.NewReader(string(text)))
			if err != nil {
				log.Fatal(err)
			}
			titleNode, err := Title(doc)
			if err != nil {
				log.Fatal(err)
			}
			title := renderNode(titleNode)
			vals = append(vals, title)
			tag := getElementById(doc, "content")
			output := renderNode(tag)
			vals = append(vals, output)
		}
	}
	c, err := os.Create("content.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	for _, s := range vals {
		c.WriteString(s)
	}
}
