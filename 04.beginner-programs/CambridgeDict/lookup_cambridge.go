package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	baseUrl := "https://dictionary.cambridge.org/dictionary/english/"
	search := "go"
	if len(os.Args) > 1 {
		search = os.Args[1]
	}
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		log.Fatal(err)
	}
	wordDb := Newsqlitedb(db)
	gotdef, err := wordDb.GetByWord(search)
	if err != nil {
		if err == ErrNotExists {
			uri := baseUrl + search
			content := curl(uri)
			if content != "" {
				_, err = wordDb.Create(search, (search + ":\n" + content))
				gotdef = content
			} else {
				gotdef = search + " not found"
			}
		}
	}
	lessPrint(gotdef)
}
