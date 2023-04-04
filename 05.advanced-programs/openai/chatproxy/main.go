package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

const userkey = "user"
const userrole = "role"

var secret = []byte("NShJJ6paNtiRSYne")
var token string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error getting env variables: %s", err)
	}
	token = os.Getenv("token")
	if len(token) == 0 {
		log.Fatal("Token not found")
	}
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
	r.GET("/api/users", withLogin(func(c *gin.Context) {
		users, _ := getUsers(db)
		session := sessions.Default(c)
		role := session.Get(userrole)
		c.JSON(http.StatusOK, gin.H{"users": users, "role": role})
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
	r.GET("/contents", withLogin(getContentHandler))
	r.POST("/contents", withLogin(searchContentHandler))
	r.GET("/contentdel", withLogin(delContentHandler))
	r.Run(":8081")
}
