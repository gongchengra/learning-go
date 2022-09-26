package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

type chapter struct {
	id      int
	title   string
	link    string
	params  string
	content string
}

func curl(chap *chapter, wg *sync.WaitGroup) {
	defer wg.Done()
	uri := "https://www.zcxswang.cc/e/extend/body/nrpost.php"
	//     payload := strings.NewReader("bid=3504&zid=860413&fid=7280539ad152c2ec6404c898f0fd940e59677691")
	payload := strings.NewReader(chap.params)
	req, _ := http.NewRequest("POST", uri, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var i interface{}
	json.Unmarshal(body, &i)
	//     fmt.Println(strings.Replace(fmt.Sprint(i), "<p>", "\n", -1))
	r := strings.NewReplacer("<p>", "\n", "</p>", "\n")
	txt := r.Replace(fmt.Sprint(i))
	re := regexp.MustCompile(`<[^>]*>`)
	chap.content = re.ReplaceAllString(txt, "")
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

func main() {
	//     resp, err := http.Get("https://www.zcxswang.cc/gudianwenxue/3504/")
	uri := "https://www.zcxswang.cc/gudianwenxue/3504/"
	if len(os.Args) > 1 {
		uri = os.Args[1]
	}
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	/*
	   body, err := ioutil.ReadAll(resp.Body)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Println(string(body))
	*/
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	docTitle := doc.Find("title").Text()
	filename := strings.Split(docTitle, "_")[0]
	// create the file
	// use strToRuneSumString so that there is a number in the front of filename
	// and you can type the filename by auto completion
	f, err := os.Create(strToRuneSumString(filename) + filename + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println(filename)
	f.WriteString(filename + "\n")
	refs := doc.Find("#booklist li a").Map(func(_ int, tag *goquery.Selection) string {
		link, _ := tag.Attr("href")
		linkText := tag.Text()
		return fmt.Sprintf("%s,%s", linkText, link)
	})
	novel := make([]chapter, len(refs))
	for i := 0; i < len(refs); i++ {
		v := strings.Split(refs[i], ",")
		title := v[0]
		link := v[1]
		params := strings.Split(link, "/")
		val := url.Values{}
		val.Add("bid", params[2])
		val.Add("zid", strings.Split(params[3], ".")[0])
		val.Add("fid", "7280539ad152c2ec6404c898f0fd940e59677691")
		novel[i].id = i
		novel[i].title = title
		novel[i].link = link
		novel[i].params = val.Encode()
	}
	var wg sync.WaitGroup
	for i := range novel {
		wg.Add(1)
		go curl(&novel[i], &wg)
	}
	wg.Wait()
	for _, chap := range novel {
		f.WriteString("\n" + chap.title + "\n")
		f.WriteString("\n" + chap.content + "\n")
	}
}
