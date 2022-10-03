package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func main() {
	router := bunrouter.New(
		bunrouter.WithNotFoundHandler(func(w http.ResponseWriter, req bunrouter.Request) error {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 - failed to find %s", req.URL.Path)
			return nil
		}),
		bunrouter.WithMethodNotAllowedHandler(func(w http.ResponseWriter, req bunrouter.Request) error {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "405 - method not allowed")
			return nil
		}),
	)
	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		w.Write([]byte("index"))
		return nil
	})
	router.POST("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusUnprocessableEntity)
		}
		fmt.Fprintln(w, req.Form)
		return nil
	})
	router.GET("/hello", func(w http.ResponseWriter, req bunrouter.Request) error {
		fmt.Fprintf(w, "hello")
		return nil
	})
	router.WithGroup("/api", func(g *bunrouter.Group) {
		g.GET("/users/:id", debugHandler)
		g.GET("/users/current", debugHandler)
		g.GET("/users/*path", debugHandler)
	})
	router.WithGroup("/api/", func(group *bunrouter.Group) {
		group.GET("/index", func(w http.ResponseWriter, req bunrouter.Request) error {
			fmt.Fprintln(w, "api/index1")
			return nil
		})
		group.GET("/hello", func(w http.ResponseWriter, req bunrouter.Request) error {
			fmt.Fprintln(w, "api/hello1")
			return nil
		})
	})
	router.WithGroup("/api2/", func(group *bunrouter.Group) {
		group.GET("/index", func(w http.ResponseWriter, req bunrouter.Request) error {
			fmt.Fprintln(w, "index2")
			return nil
		})
		group.GET("/hello", func(w http.ResponseWriter, req bunrouter.Request) error {
			fmt.Fprintln(w, "hello2")
			return nil
		})
	})
	log.Println("listening on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", router))
}

func debugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
	})
}
