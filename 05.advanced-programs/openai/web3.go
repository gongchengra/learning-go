package main

import (
	"fmt"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	User{ID: 1, Username: "alice", Password: "$2a$10$nCClD0F8waELXkghQ8QKQ.gTA0nwh70sGkNl3Iuc6MrSStfsHn6hW"}, // Password is "123456" for test purpose
	User{ID: 2, Username: "bob", Password: "$2a$10$NShJJ6paNtiRSYne/SAILO.sbyEwPmk5RsZXp/JX/1eIYR.fokNda"},   // Password is "654321"
	User{ID: 3, Username: "alan", Password: "$2a$10$EKCW1rrZnVC5dnhc6jJbBOroOl20yz5SLybrFw12nXy0igY08gY7i"},
}
var secret = []byte("NShJJ6paNtiRSYne")

const userkey = "user"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

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
	router.POST("/", func(c *gin.Context) {
		session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("password")
		for _, user := range Users {
			if user.Username == username {
				if CheckPassword(user.Password, password) {
					session.Set(userkey, user.ID)
					if err := session.Save(); err != nil {
						c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{"error": "Failed to save session"})
						return
					}
					c.HTML(http.StatusOK, "input.tmpl", gin.H{})
				} else {
					c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid credentials"})
				}
				return
			}
		}
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid user"})
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
	})
	router.GET("/logout", logout)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	// Start server
	router.Run(":8081")
}
