package main

// https://zetcode.com/golang/pipe/
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type PayLoad struct {
	Content string
}

func main() {
	r, w := io.Pipe()
	go func() {
		defer w.Close()
		err := json.NewEncoder(w).Encode(&PayLoad{Content: "Hello there!"})
		if err != nil {
			log.Fatal(err)
		}
	}()
	resp, err := http.Post("https://httpbin.org/post", "application/json", r)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
