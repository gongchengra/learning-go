package main

// https://zetcode.com/golang/servemux/
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func main() {
	mux := http.NewServeMux()
	now := time.Now()
	mux.HandleFunc("/today", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte(now.Format(time.ANSIC)))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello there!")
	})
	hello := &helloHandler{}
	mux.Handle("/hello", hello)
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
