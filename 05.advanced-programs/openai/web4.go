package main

import (
	"database/sql"
	"fmt"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

var Users = []User{
	User{ID: 1, Username: "alice", Password: "$2a$10$nCClD0F8waELXkghQ8QKQ.gTA0nwh70sGkNl3Iuc6MrSStfsHn6hW"}, // Password is "123456" for test purpose
	User{ID: 2, Username: "bob", Password: "$2a$10$NShJJ6paNtiRSYne/SAILO.sbyEwPmk5RsZXp/JX/1eIYR.fokNda"},   // Password is "654321"
	User{ID: 3, Username: "alan", Password: "$2a$10$EKCW1rrZnVC5dnhc6jJbBOroOl20yz5SLybrFw12nXy0igY08gY7i"},
}
var secret = []byte("NShJJ6paNtiRSYne")

const userkey = "user"

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func chat(input string) (output string) {
	token := os.Getenv("token")
	if len(token) == 0 {
		return
	}
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
		fmt.Println(err)
		panic(err)
	}
	output = resp.Choices[0].Message.Content
	return output
}

func checkLogin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	}
	c.Next()
}

func findUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		id := c.Query("id")
		if id == "" {
			id = strconv.Itoa(user.(int))
		}
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, username FROM users where id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		var foundUser User
		for rows.Next() {
			err := rows.Scan(&foundUser.ID, &foundUser.Username)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.HTML(http.StatusOK, "user.tmpl", gin.H{"user": foundUser})
	}
}

func getUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, username, password FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Password)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "user.tmpl", gin.H{"users": users})
	}
}

func addUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var role int
		db.QueryRow("SELECT role_id FROM users where id = ?", user).Scan(&role)
		log.Println("role id", user, role)
		if role != 1 {
			c.HTML(http.StatusNotFound, "user.tmpl", gin.H{"message": "Please login in as admin"})
			return
		}
		username := c.PostForm("username")
		password := c.PostForm("password")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		statement, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()
		_, err = statement.Exec(username, hashedPassword)
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusSeeOther, "/users")
	}
}

func delUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		id := c.Query("id")
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var role int
		db.QueryRow("SELECT role_id FROM users where id = ?", user).Scan(&role)
		if role != 1 {
			c.HTML(http.StatusNotFound, "user.tmpl", gin.H{"message": "Please login in as admin"})
			return
		}
		stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(id)
		if err != nil {
			log.Fatal(err)
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		if rowsAffected == 0 {
			c.HTML(http.StatusNotFound, "user.tmpl", gin.H{"message": "User not found"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/users")
	}
}

func getContent(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("select id, prompt, answer, userid from contents")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var contents []Content
		for rows.Next() {
			var content Content
			err := rows.Scan(&content.ID, &content.Prompt, &content.Answer, &content.UserID)
			if err != nil {
				log.Fatal(err)
			}
			contents = append(contents, content)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "content.tmpl", gin.H{"contents": contents})
	}
}

func delContent(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		c.Abort()
		return
	} else {
		id := c.Query("id")
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var role int
		db.QueryRow("SELECT role_id FROM users where id = ?", user).Scan(&role)
		log.Println("role id", user, role)
		if role != 1 {
			c.HTML(http.StatusUnauthorized, "content.tmpl", gin.H{"error": "Please login in as admin"})
			return
		}
		stmt, err := db.Prepare("DELETE FROM contents WHERE id = ?")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(id)
		if err != nil {
			log.Fatal(err)
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		if rowsAffected == 0 {
			c.HTML(http.StatusNotFound, "content.tmpl", gin.H{"message": "Content not found"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/contents")
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{"error": "Failed to save session"})
		return
	}
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"message": "Successfully logged out"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{})
		} else {
			c.HTML(http.StatusOK, "input.tmpl", gin.H{})
		}
	})
	// Authenticate user
	router.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("password")
		password = strings.TrimSpace(password) // 去掉换行符和空格
		db, err := sql.Open("sqlite3", "./chat.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		query := "SELECT id, username, password FROM users WHERE username=?"
		var id int
		var dbUsername string
		var dbPassword string
		err = db.QueryRow(query, username).Scan(&id, &dbUsername, &dbPassword)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid user"})
			return
		}
		fmt.Println(dbPassword, password)
		if CheckPassword(dbPassword, password) {
			session.Set(userkey, id)
			if err := session.Save(); err != nil {
				c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{"error": "Failed to save session"})
				return
			}
			c.HTML(http.StatusOK, "input.tmpl", gin.H{})
		} else {
			c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid credentials"})
		}
	})
	router.GET("/input", func(c *gin.Context) {
		checkLogin(c)
		c.HTML(http.StatusOK, "input.tmpl", gin.H{})
	})
	router.POST("/input", func(c *gin.Context) {
		checkLogin(c)
		input := c.PostForm("input")
		output := chat(input)
		c.HTML(http.StatusOK, "input.tmpl", gin.H{
			"input":  input,
			"output": output,
		})
		if output != "" {
			db, err := sql.Open("sqlite3", "./chat.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			session := sessions.Default(c)
			user := session.Get(userkey)
			statement, err := db.Prepare("INSERT INTO contents (prompt, answer, userid) VALUES (?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer statement.Close()
			_, err = statement.Exec(input, output, user)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	router.GET("/users", getUser)
	router.GET("/user", findUser)
	router.POST("/useradd", addUser)
	router.GET("/userdel", delUser)
	router.GET("/contents", getContent)
	router.GET("/contentdel", delContent)
	router.GET("/logout", logout)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	// Start server
	router.Run(":8080")
}
