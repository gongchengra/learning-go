package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func curl(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}
	//     def := doc.Find(".ddef_d,.examp").Text()
	text := doc.Find(".ddef_d,.dexamp").Map(func(_ int, sel *goquery.Selection) string {
		if sel.HasClass("ddef_d") {
			return fmt.Sprintf("%s%s", "\nDef: ", strings.Trim(sel.Text(), ": "))
		} else {
			return fmt.Sprintf("%s%s", "Example: ", strings.TrimSpace(sel.Text()))
		}
	})
	return strings.Join(text, "\n")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func lessPrint(c string) {
	r := strings.Split(c, "\n")
	var input int
	for i := range r {
		if i > 0 && i%20 == 0 {
			fmt.Scanf("%d", &input)
		}
		fmt.Println(i, r[i])
	}
}
