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
	ID      int    `json:"id"`
	Prompt  string `json:"prompt"`
	Answer  string `json:"answer"`
	UserID  int    `json:"userid"`
	IsImage int    `json:"isImage"`
}

func getContent(db *sql.DB) ([]Content, error) {
	rows, err := db.Query("SELECT id, prompt, answer, userid, isImage FROM contents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var contents []Content
	for rows.Next() {
		var content Content
		err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID, &content.IsImage)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, rows.Err()
}

func getLastContent(db *sql.DB, userID int) (Content, error) {
	row := db.QueryRow("SELECT id, prompt, answer, userid FROM contents WHERE userid=? AND is_deleted=0 ORDER BY id DESC LIMIT 1", userID)
	var content Content
	err := row.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID, &content.IsImage)
	if err != nil {
		return Content{}, err
	}
	return content, nil
}

func getPageOfContent(db *sql.DB, user int, page int, pageSize int) ([]Content, bool) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, prompt, answer, userid, isImage FROM contents where is_deleted = 0 and userid = %d ORDER BY id DESC LIMIT %d OFFSET %d", user, pageSize, offset)
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
		err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID, &content.IsImage)
		if err != nil {
			return nil, false
		}
		contents = append(contents, content)
	}
	queryCount := fmt.Sprintf("SELECT count(*) FROM contents where is_deleted = 0 and userid = %d", user)
	db.QueryRow(queryCount).Scan(&totalRows)
	hasNextPage := (totalRows > offset+len(contents))
	return contents, hasNextPage
}

func addContent(db *sql.DB, input, output string, userID int, isImage int) error {
	statement, err := db.Prepare("INSERT INTO contents (prompt, answer, userid, isImage) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(input, output, userID, isImage)
	return err
}

func checkExist(db *sql.DB, input string) (Content, error) {
	row := db.QueryRow("SELECT id, prompt, answer, userid FROM contents WHERE prompt=?  ORDER BY id DESC LIMIT 1", input)
	var content Content
	err := row.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID)
	if err != nil {
		return Content{}, err
	}
	return content, nil
}

func addContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	userID := session.Get(userkey).(int)
	input := c.PostForm("input")
	existContent, err := checkExist(db, input)
	if err == nil {
		c.HTML(http.StatusOK, "input.tmpl", gin.H{
			"title":  "Input Prompt",
			"input":  input,
			"output": existContent.Answer,
		})
		return
	} else {
		log.Println(err)
	}
	checked := false
	continueVal := c.PostForm("continue")
	if continueVal == "yes" {
		checked = true
	}
	context := ""
	if checked {
		lastContent, err := getLastContent(db, userID)
		if err == nil {
			context = lastContent.Prompt + lastContent.Answer
		}
	}
	output := chat(input, context)
	c.HTML(http.StatusOK, "input.tmpl", gin.H{
		"title":  "Input Prompt",
		"input":  input,
		"output": output,
	})
	if context != "" {
		input = context + input
	}
	err = addContent(db, input, output, userID, 0)
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
	rows, err := db.Query("SELECT id, prompt, answer, isImage FROM contents WHERE is_deleted = 0 and (prompt LIKE '%'||?||'%' OR answer LIKE '%'||?||'%') ORDER BY id DESC", search, search)
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

func getPageOfSearchResults(db *sql.DB, user int, search string, page int, pageSize int) ([]Content, bool, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, prompt, answer, userid, isImage FROM contents where is_deleted = 0 and userid = %d and (prompt LIKE '%%%s%%' OR answer LIKE '%%%s%%') LIMIT %d OFFSET %d", user, search, search, pageSize, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	var (
		contents  []Content
		totalRows int
	)
	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID, &content.IsImage); err != nil {
			return nil, false, err
		}
		contents = append(contents, content)
	}
	queryCount := fmt.Sprintf("SELECT count(*) FROM contents where is_deleted = 0 and userid = %d and (prompt LIKE '%%%s%%' OR answer LIKE '%%%s%%')", user, search, search)
	if err := db.QueryRow(queryCount).Scan(&totalRows); err != nil {
		return nil, false, err
	}
	hasNextPage := (totalRows > offset+len(contents))
	return contents, hasNextPage, nil
}

func contentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	user := session.Get(userkey).(int)
	if c.Request.Method == "GET" {
		search := c.Query("search")
		if search != "" {
			page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
			if err != nil {
				c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to parse page number"})
				return
			}
			pagesize, err := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
			if err != nil {
				c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to parse page size"})
				return
			}
			contents, hasNextPage, err := getPageOfSearchResults(db, user, search, page, pagesize)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to search content"})
				return
			}
			prevPage := ""
			if page > 1 {
				prevPage = fmt.Sprintf("/contents/search?search=%s&page=%d&pagesize=%d", search, page-1, pagesize)
			}
			nextPage := ""
			if hasNextPage {
				nextPage = fmt.Sprintf("/contents/search?search=%s&page=%d&pagesize=%d", search, page+1, pagesize)
			}
			c.HTML(http.StatusOK, "content.tmpl", gin.H{
				"title":    "Contents",
				"contents": contents,
				"prevPage": prevPage,
				"nextPage": nextPage,
				"search":   search,
			})
		} else {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
			contents, hasNextPage := getPageOfContent(db, user, page, pagesize)
			prevPage := ""
			if page > 1 {
				prevPage = fmt.Sprintf("/contents?page=%d&pagesize=%d", page-1, pagesize)
			}
			nextPage := ""
			if hasNextPage {
				nextPage = fmt.Sprintf("/contents?page=%d&pagesize=%d", page+1, pagesize)
			}
			c.HTML(http.StatusOK, "content.tmpl", gin.H{
				"title":    "Contents",
				"contents": contents,
				"prevPage": prevPage,
				"nextPage": nextPage,
			})
		}
	} else if c.Request.Method == "POST" {
		search := c.PostForm("search")
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to parse page number"})
			return
		}
		pagesize, err := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to parse page size"})
			return
		}
		contents, hasNextPage, err := getPageOfSearchResults(db, user, search, page, pagesize)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to search content"})
			return
		}
		prevPage := ""
		if page > 1 {
			prevPage = fmt.Sprintf("/contents/search?search=%s&page=%d&pagesize=%d", search, page-1, pagesize)
		}
		nextPage := ""
		if hasNextPage {
			nextPage = fmt.Sprintf("/contents/search?search=%s&page=%d&pagesize=%d", search, page+1, pagesize)
		}
		c.HTML(http.StatusOK, "content.tmpl", gin.H{
			"title":    "Contents",
			"contents": contents,
			"prevPage": prevPage,
			"nextPage": nextPage,
			"search":   search,
		})
	}
}

func apiContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	user := session.Get(userkey).(int)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
	if c.Request.Method == "GET" {
		search := c.Query("search")
		if search != "" {
			contents, _, err := getPageOfSearchResults(db, user, search, page, pagesize)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to search content",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"contents": contents,
				"search":   search,
			})
		} else {
			contents, _ := getPageOfContent(db, user, page, pagesize)
			c.JSON(http.StatusOK, gin.H{
				"contents": contents,
			})
		}
	} else if c.Request.Method == "POST" {
		search := c.PostForm("search")
		contents, _, err := getPageOfSearchResults(db, user, search, page, pagesize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to search content",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"contents": contents,
			"search":   search,
		})
	}
}
