package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

const userkey = "user"
const userrole = "role"

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

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "input.tmpl", gin.H{})
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
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
	r.GET("/account", withLogin(func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(userkey)
		id := user.(int)
		foundUser, _ := findUser(db, id)
		c.HTML(http.StatusOK, "account.tmpl", gin.H{"user": foundUser})
	}))
	r.POST("/account", withLogin(updateUserPassword))
	r.POST("/useradd", withLogin(register))
	r.GET("/userdel", withLogin(delUserHandler))
	r.GET("/logout", logout)
	r.GET("/contents", withLogin(getContentHandler))
	r.POST("/contents", withLogin(searchContentHandler))
	r.GET("/contentdel", withLogin(delContentHandler))
	r.Run(":8081")
}
