package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func findUser(db *sql.DB, id int) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, username FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Username)
	return user, err
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username)
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
