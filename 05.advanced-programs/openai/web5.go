package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const userkey = "user"
const userrole = "role"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Content struct {
	ID     int    `json:"id"`
	Prompt string `json:"username"`
	Answer string `json:"password"`
	UserID int    `json:"userid"`
}

var secret []byte
var token string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error getting env variables: %s", err)
	}
	token = os.Getenv("token")
	secret = []byte(os.Getenv("secret"))
	if len(token) == 0 {
		log.Fatal("Token not found")
	}
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func chat(input string) (output string) {
	c := gpt35.NewClient(token)
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: input,
			},
		},
	}
	resp, err := c.GetChat(req)
	if err != nil {
		log.Println(err)
		return
	}
	output = resp.Choices[0].Message.Content
	return output
}

func withLogin(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
			c.Abort()
			return
		}
		handler(c)
	}
}

func findUser(db *sql.DB, id int) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Username)
	return user, err
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func addUser(db *sql.DB, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(username, hashedPassword)
	return err
}

func delUser(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
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
		return fmt.Errorf("User not found")
	}
	return nil
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

func getPageOfContent(db *sql.DB, page int, pageSize int) ([]Content, bool) {
	if page < 1 {
		page = 0
	}
	if pageSize < 1 {
		pageSize = 5
	}
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, prompt, answer, userid FROM contents LIMIT %d OFFSET %d", pageSize, offset)
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
	db.QueryRow("SELECT count(*) FROM contents").Scan(&totalRows)
	hasNextPage := (totalRows > offset+len(contents))
	return contents, hasNextPage
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

func login(c *gin.Context) {
	session := sessions.Default(c)
	db := c.MustGet("db").(*sql.DB)
	username := c.PostForm("username")
	password := c.PostForm("password")
	password = strings.TrimSpace(password)
	var id int
	var dbUsername string
	var dbPassword string
	var role int
	err := db.QueryRow("SELECT id, username, password, role_id FROM users WHERE username=?", username).Scan(&id, &dbUsername, &dbPassword, &role)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid user"})
		return
	}
	if CheckPassword(dbPassword, password) {
		session.Set(userkey, id)
		session.Set(userrole, role)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{"error": "Failed to save session"})
			return
		}
		c.HTML(http.StatusOK, "input.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid password"})
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func register(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	username := c.PostForm("username")
	password := c.PostForm("password")
	password = strings.TrimSpace(password)
	if len(password) < 6 {
		c.HTML(http.StatusBadRequest, "user.tmpl", gin.H{"error": "Password must be at least 6 characters long"})
		return
	}
	var existingUser User
	row := db.QueryRow("SELECT id, username FROM users WHERE username = ?", username)
	err := row.Scan(&existingUser.ID, &existingUser.Username)
	if err == nil {
		c.HTML(http.StatusBadRequest, "user.tmpl", gin.H{"error": "User already exists"})
		return
	}
	err = addUser(db, username, password)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "user.tmpl", gin.H{"error": "Failed to add user"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/users")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "input.tmpl", gin.H{})
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

func delContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	role := session.Get(userrole)
	if role != 1 {
		c.HTML(http.StatusUnauthorized, "content.tmpl", gin.H{"error": "Please login in as admin"})
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	err := delContent(db, id)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "content.tmpl", gin.H{"error": "Failed to delete content"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/contents")
}

func main() {
	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := cookie.NewStore(secret)
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("html/*")
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", withLogin(index))
	r.POST("/login", login)
	r.GET("/input", withLogin(func(c *gin.Context) {
		c.HTML(http.StatusOK, "input.tmpl", gin.H{})
	}))
	r.POST("/input", withLogin(addContentHandler))
	r.GET("/users", withLogin(func(c *gin.Context) {
		users, _ := getUsers(db)
		session := sessions.Default(c)
		role := session.Get(userrole)
		c.HTML(http.StatusOK, "user.tmpl", gin.H{"users": users, "role": role})
	}))
	r.GET("/user", withLogin(func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		if id == 0 {
			session := sessions.Default(c)
			user := session.Get(userkey)
			id = user.(int)
		}
		foundUser, _ := findUser(db, id)
		c.HTML(http.StatusOK, "user.tmpl", gin.H{"user": foundUser})
	}))
	r.POST("/useradd", withLogin(register))
	r.GET("/userdel", withLogin(func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get(userrole)
		if role != 1 {
			c.HTML(http.StatusNotFound, "user.tmpl", gin.H{"message": errors.New("Please login in as admin")})
			return
		}
		id, _ := strconv.Atoi(c.Query("id"))
		err := delUser(db, id)
		if err != nil {
			c.HTML(http.StatusNotFound, "user.tmpl", gin.H{"message": err})
		} else {
			c.Redirect(http.StatusSeeOther, "/users")
		}
	}))
	r.GET("/logout", logout)
	r.GET("/contents", withLogin(func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
		pagesize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "5"))
		contents, hasNextPage := getPageOfContent(db, page, pagesize)
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
	}))
	r.GET("/contentdel", withLogin(delContentHandler))
	r.Run(":8080")
}

func nl2br(str string) template.HTML {
	return template.HTML(strings.ReplaceAll(str, "\n", "<br>"))
}
