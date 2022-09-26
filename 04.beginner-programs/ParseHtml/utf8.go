package main

// go build utf8.go util.go
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return ""
	}
	return buf.String()
}

func main() {
	f, err := os.Open(".")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	if dirs, err := f.Readdir(-1); err == nil {
		SortFileNameAscend(dirs)
		for _, d := range dirs {
			filename := d.Name()
			if strings.Contains(filename, "txt") {
				fmt.Println(filename)
				htmlname := strings.Replace(filename, "txt", "html", 1)
				fmt.Println(htmlname)
				text, err := os.ReadFile(htmlname)
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
				fmt.Println(title)
				content, err := ioutil.ReadFile(filename)
				if err != nil {
					fmt.Println(err)
				}
				var i interface{}
				json.Unmarshal(content, &i)
				fmt.Println(err, i)
			}
		}
	}
}
