package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Content struct {
	ID     int    `json:"id"`
	Prompt string `json:"username"`
	Answer string `json:"password"`
	UserID int    `json:"userid"`
}

func getContent(db *sql.DB) ([]Content, error) {
	rows, err := db.Query("SELECT id, prompt, answer, userid FROM contents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var contents []Content
	for rows.Next() {
		var content Content
		err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, rows.Err()
}

func getPageOfContent(db *sql.DB, user int, page int, pageSize int) ([]Content, bool) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, prompt, answer, userid FROM contents where is_deleted = 0 and userid = %d LIMIT %d OFFSET %d", user, pageSize, offset)
	log.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		return nil, false
	}
	defer rows.Close()
	var (
		contents  []Content
		totalRows int
	)
	for rows.Next() {
		var content Content
		err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID)
		if err != nil {
			return nil, false
		}
		contents = append(contents, content)
	}
	queryCount := fmt.Sprintf("SELECT count(*) FROM contents where is_deleted = 0 and userid = %d", user)
	db.QueryRow(queryCount).Scan(&totalRows)
	log.Println(totalRows)
	hasNextPage := (totalRows > offset+len(contents))
	return contents, hasNextPage
}

func getContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
	session := sessions.Default(c)
	user := session.Get(userkey)
	id := user.(int)
	contents, hasNextPage := getPageOfContent(db, id, page, pagesize)
	prevPage := ""
	if page > 1 {
		prevPage = fmt.Sprintf("/contents?page=%d&pagesize=%d", page-1, pagesize)
	}
	nextPage := ""
	if hasNextPage {
		nextPage = fmt.Sprintf("/contents?page=%d&pagesize=%d", page+1, pagesize)
	}
	c.HTML(http.StatusOK, "content.tmpl", gin.H{
		"contents": contents,
		"prevPage": prevPage,
		"nextPage": nextPage,
	})
}

func addContent(db *sql.DB, input, output string, userID int) error {
	statement, err := db.Prepare("INSERT INTO contents (prompt, answer, userid) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(input, output, userID)
	return err
}

func addContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	userID := session.Get(userkey).(int)
	input := c.PostForm("input")
	output := chat(input)
	c.HTML(http.StatusOK, "input.tmpl", gin.H{
		"input":  input,
		"output": output,
	})
	err := addContent(db, input, output, userID)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "input.tmpl", gin.H{"error": "Failed to add content"})
		return
	}
}

func delContent(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM contents WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Content not found")
	}
	return nil
}

func softDeleteContent(db *sql.DB, id int) error {
	stmt, err := db.Prepare("UPDATE contents SET is_deleted = ? WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(1, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Content not found")
	}
	return nil
}

func delContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	role := session.Get(userrole)
	if role != 1 {
		id, _ := strconv.Atoi(c.Query("id"))
		err := softDeleteContent(db, id)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to delete content"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/contents")
	} else {
		id, _ := strconv.Atoi(c.Query("id"))
		err := delContent(db, id)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to delete content"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/contents")
	}
}

func searchContent(db *sql.DB, search string) ([]Content, error) {
	var contents []Content
	rows, err := db.Query("SELECT id, prompt, answer FROM contents WHERE is_deleted = 0 and (prompt LIKE '%'||?||'%' OR answer LIKE '%'||?||'%')", search, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ID, &content.Prompt, &content.Answer); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return contents, nil
}

func searchContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	_ = session.Get(userkey).(int)
	search := c.PostForm("search")
	contents, err := searchContent(db, search)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to search content"})
		return
	}
	c.HTML(http.StatusOK, "content.tmpl", gin.H{
		"contents": contents,
	})
}
