package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++
	fmt.Fprintln(w, "Counter:", ct.counter)
}

func main() {
	th := &CounterHandler{counter: 0}
	http.Handle("/count", th)
	http.ListenAndServe(":8080", nil)
}
