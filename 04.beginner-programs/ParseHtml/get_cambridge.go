package main

// go build get_tianya.go util.go
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func curl(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	readFile, err := os.Open("words.log")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		word := fileScanner.Text()
		if !fileExists(word + ".txt") {
			uri := baseUrl + word
			f, err := os.Create(word + ".txt")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			f.WriteString(word + ":\n")
			f.WriteString(curl(uri))
			time.Sleep(1 * time.Second)
		}
	}
}
