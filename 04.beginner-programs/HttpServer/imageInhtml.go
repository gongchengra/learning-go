package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./data"))
	http.Handle("/data/", http.StripPrefix("/data/", fs))
	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "image.html")
	})
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
