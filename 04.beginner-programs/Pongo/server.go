package main

import (
	"github.com/flosch/pongo2/v5"
	"net/http"
)

type User struct {
	Name       string
	Occupation string
}

var tpl = pongo2.Must(pongo2.FromFile("users.html"))

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "John Doe", Occupation: "gardener"},
		{Name: "Roger Roe", Occupation: "driver"},
		{Name: "Peter Smith", Occupation: "teacher"},
	}
	err := tpl.ExecuteWriter(pongo2.Context{"users": users}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
