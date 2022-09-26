package main

// https://gist.github.com/gongchengra/0535db596cbdba601ffcf3752d22f2d5
import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"unicode/utf8"
)

func Decodegbk(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func strToRuneSumString(s string) string {
	res := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		res += int(r)
		s = s[size:]
	}
	return strconv.Itoa(int(res))
}

func SortFileNameAscend(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
}

func Title(doc *html.Node) (*html.Node, error) {
	var title *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" {
			title = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if title != nil {
		return title, nil
	}
	return nil, errors.New("Missing <title> in the node tree")
}
