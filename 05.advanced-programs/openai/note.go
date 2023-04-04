package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Note struct {
	ID  int    `json:"id"`
	Txt string `json:"text"`
}

var db *sql.DB // database connection

func openDB() {
	var err error
	db, err = sql.Open("sqlite3", "./note.db")
	if err != nil {
		log.Fatal(err)
	}
}

func getNote(c *gin.Context) {
	rows, err := db.Query("SELECT id, note FROM notes where is_deleted = 0")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Txt)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "/note.tmpl", gin.H{
		"notes": notes,
	})
}

func getNoteAll(c *gin.Context) {
	rows, err := db.Query("SELECT id, note FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Txt)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "/notes.tmpl", gin.H{
		"notes": notes,
	})
}

func addNote(c *gin.Context) {
	note := c.PostForm("note")
	statement, err := db.Prepare("INSERT INTO notes (note) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	if _, err := statement.Exec(note); err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func delNote(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		c.HTML(http.StatusNotFound, "/note.tmpl", gin.H{"message": "Note not found"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func softDeleteNote(c *gin.Context) {
	id := c.Query("id")
	stmt, err := db.Prepare("UPDATE notes SET is_deleted = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(1, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		c.HTML(http.StatusNotFound, "/note.tmpl", gin.H{"message": "Note not found"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func main() {
	openDB()
	defer db.Close()
	gin.SetMode(gin.ReleaseMode)
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.SetHTMLTemplate(t)
	router.GET("/", getNote)
	router.GET("/all", getNoteAll)
	router.POST("/", addNote)
	router.GET("/notedel", softDeleteNote)
	router.GET("/del", delNote)
	defer db.Close()
	router.Run(":8080")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
