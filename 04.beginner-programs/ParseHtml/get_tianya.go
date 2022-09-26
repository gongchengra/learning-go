package main

// go build get_tianya.go util.go
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
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
	b, err := Decodegbk(body)
	//     fmt.Println(string(b))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	content := doc.Find("#htmlContent").Text()
	return content
}

func main() {
	//     resp, err := http.Get("https://www.zcxswang.cc/gudianwenxue/3504/")
	uri := "https://www.tianyabook.com/shu/21372.html"
	if len(os.Args) > 1 {
		uri = os.Args[1]
	}
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	b, err := Decodegbk(body)
	//     fmt.Println(string(b))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}
	docTitle := doc.Find("title").Text()
	//     fmt.Println(docTitle)
	filename := strings.Split(docTitle, "_")[0]
	//     filename := docTitle
	// create the file
	// use strToRuneSumString so that there is a number in the front of filename
	// and you can type the filename by auto completion
	f, err := os.Create(strToRuneSumString(filename) + filename + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer f.Close()
	fmt.Println(filename)
	f.WriteString(filename + "\n")
	refs := doc.Find("#list-chapterAll dl dd a").Map(func(_ int, tag *goquery.Selection) string {
		link, _ := tag.Attr("href")
		linkText := tag.Text()
		return fmt.Sprintf("%s,%s", linkText, link)
	})
	//     fmt.Println(refs)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(refs); i++ {
		v := strings.Split(refs[i], ",")
		title := v[0]
		//         link := v[1]
		link := "https://www.tianyabook.com" + v[1]
		fmt.Println(title, link)
		f.WriteString("\n" + title + "\n")
		f.WriteString(curl(link))
		//         n := time.Duration(5 + rand.Intn(10))
		//                 time.Sleep(n * time.Second)
	}
}
