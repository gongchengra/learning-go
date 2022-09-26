package main

// https://zetcode.com/golang/http-client/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("https://httpbin.org/post",
		url.Values{"name": {"John Doe"}, "message": {"Hey!"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
