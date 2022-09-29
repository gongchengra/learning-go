package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

// Compile templates on start of the application
var index = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Lastpass Search</title>
  </head>
  <body>
    <form enctype="multipart/form-data" action="/" method="post">
	  <label for="search">Word:</label>
      <input type="text" name="search" />
      <input type="submit" value="search" />
    </form>
	<div>{{.}}</div>
  </body>
</html>
`))

func divPrint(c string) string {
	r := strings.Split(c, "\n")
	res := ""
	for i := range r {
		res += "<div>" + r[i] + "</div>"
	}
	return res
}

const fileName = "words.db"

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	wordDb := Newsqlitedb(db)
	r := mux.NewRouter()
	r.HandleFunc("/search/{word}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]
		fmt.Println("searched ", word)
		switch r.Method {
		case "GET":
			gotdef, err := wordDb.GetByWord(word)
			if err == ErrNotExists {
				w.Write(nil)
			} else {
				w.Write([]byte(gotdef))
			}
		case "POST":
			def := r.FormValue("def")
			if def != "" {
				_, err = wordDb.Create(word, def)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			index.ExecuteTemplate(w, "index.html", nil)
		case "POST":
			search := r.FormValue("search")
			if strings.TrimSpace(search) == "" {
				return
			}
			gotdef, err := wordDb.GetByWord(search)
			if err != nil {
				if err == ErrNotExists {
					uri := baseUrl + search
					content := curl(uri)
					if content != "" {
						_, err = wordDb.Create(search, (search + ":\n" + content))
						if err != nil {
							log.Fatal(err)
						}
						gotdef = content
					} else {
						gotdef = search + " not found"
					}
				} else {
					log.Fatal(err)
				}
			}
			index.ExecuteTemplate(w, "index.html", divPrint(gotdef))
		}
	})
	http.ListenAndServe(":8080", r)
}
