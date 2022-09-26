package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	uri := "https://www.google.com/search?q=%E6%9C%80%E5%A5%BD%E7%9A%84%E7%BD%91%E7%BB%9C%E5%B0%8F%E8%AF%B4"
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
	fmt.Println(string(body))
}
