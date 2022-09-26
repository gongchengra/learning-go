package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://webcode.me")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resp.Header {
		fmt.Printf("%s %s\n", k, v)
	}
}
