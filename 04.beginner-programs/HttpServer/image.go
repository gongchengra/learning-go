package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/image", handler)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	buf, err := os.ReadFile("sid.png")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/png")
	// serve as attachment
	w.Header().Set("Content-Disposition", `attachment;filename="sid.png"`)
	w.Write(buf)
}
