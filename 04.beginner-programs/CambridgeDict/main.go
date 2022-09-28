package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"
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
	wordRepository := NewSQLiteRepository(db)
	// Index route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			index.ExecuteTemplate(w, "index.html", nil)
		case "POST":
			search := r.FormValue("search")
			if strings.TrimSpace(search) == "" {
				return
			}
			gotdef, err := wordRepository.GetByWord(search)
			if err != nil {
				if err == ErrNotExists {
					uri := baseUrl + search
					content := curl(uri)
					if content != "" {
						_, err = wordRepository.Create(search, string(content))
						if err != nil {
							log.Fatal(err)
						}
						gotdef = content
					} else {
						gotdef = search + " not found"
					}
				}
			}
			index.ExecuteTemplate(w, "index.html", divPrint(gotdef))
		}

	})
	http.ListenAndServe(":8080", nil)
}
