package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		if name == "" {
			name = "guest"
		}
		fmt.Fprintf(resp, "Hello %s!", name)
	})
	r.HandleFunc("/hi/{name}", func(resp http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]
		fmt.Fprintf(resp, "Hi %s!", name)
	})
	r.HandleFunc("/now", func(resp http.ResponseWriter, _ *http.Request) {
		now := time.Now()
		payload := make(map[string]string)
		payload["now"] = now.Format(time.ANSIC)
		payload["name"] = "alan"
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(payload)
	})
	s1 := r.PathPrefix("/path1").Subrouter()
	s1.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(resp, "Subroute 1")
	})
	s2 := r.PathPrefix("/path2").Subrouter()
	s2.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(resp, "Subroute 2")
	})
	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}
