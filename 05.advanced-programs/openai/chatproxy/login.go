package main

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"time"
)

var rateLimit = make(map[string]int)

// Check if the IP address has exceeded the login attempt limit

func isRateLimited(ip string) bool {
	count, ok := rateLimit[ip]
	if ok && count >= 5 {
		return true
	}
	return false
}

// Increment the login attempt counter for the IP address

func incrementRateLimit(ip string) {
	rateLimit[ip]++
	// Reset the counter after 5 minutes
	time.AfterFunc(5*time.Minute, func() {
		rateLimit[ip] = 0
	})
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	db := c.MustGet("db").(*sql.DB)
	username := c.PostForm("username")
	password := c.PostForm("password")
	password = strings.TrimSpace(password)
	ip := c.ClientIP()
	if isRateLimited(ip) {
		c.HTML(http.StatusTooManyRequests, "login.tmpl", gin.H{"error": "Too many login attempts. Please try again later."})
		return
	}
	var id int
	var dbUsername string
	var dbPassword string
	var role int
	err := db.QueryRow("SELECT id, username, password, role_id FROM users WHERE username=?", username).Scan(&id, &dbUsername, &dbPassword, &role)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"error": "Invalid user"})
		incrementRateLimit(ip)
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
		incrementRateLimit(ip)
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, "/")
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
