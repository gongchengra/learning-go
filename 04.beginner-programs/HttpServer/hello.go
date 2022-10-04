package main

// refers to https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, there\n")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about page")
	})
	http.HandleFunc("/news", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "news page")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			w.Write([]byte("404 - not found\n"))
			return
		}
		fmt.Fprintln(w, "home page")
	})
	http.HandleFunc("/ua", func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		fmt.Fprintf(w, "User agent: %s\n", ua)
	})
	http.HandleFunc("/path/", PathHandler)
	// curl localhost:8080/query/?name=Peter
	http.HandleFunc("/query/", QueryHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	name := "guest"
	if ok {
		name = keys[0]
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
