package main

import (
	"fmt"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
