package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("date")
	pr, pw := io.Pipe()
	defer pw.Close()
	cmd.Stdout = pw
	cmd.Stderr = pw
	go io.Copy(w, pr)
	cmd.Run()
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
