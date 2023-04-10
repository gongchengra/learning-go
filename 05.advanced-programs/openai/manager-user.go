package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func main() {
	fmt.Print("Enter database path: ")
	var dbPath string
	fmt.Scanln(&dbPath)
	if dbPath == "" {
		dbPath = "chat.db"
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()
	var choice string
	for {
		fmt.Println("1. List all users")
		fmt.Println("2. Add a new user")
		fmt.Println("3. Delete a user")
		fmt.Println("4. Change a user's password")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			listUsers(db)
		case "2":
			addUser(db)
		case "3":
			deleteUser(db)
		case "4":
			changePassword(db)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func listUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		fmt.Println("Error querying users:", err)
		return
	}
	defer rows.Close()
	users := []User{}
	fmt.Println("Users:")
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			fmt.Println("Error scanning user:", err)
			return
		}
		users = append(users, u)
		fmt.Printf("%d. %s\n", u.ID, u.Username)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating users:", err)
		return
	}
	fmt.Println()
}

func addUser(db *sql.DB) {
	var username string
	var password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	_, err = db.Exec("INSERT INTO users(username, password, role_id) VALUES(?, ?, 2)", username, hashedPassword)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return
	}
	fmt.Println("User added successfully.")
}

func deleteUser(db *sql.DB) {
	var userID int
	fmt.Print("Enter user ID: ")
	_, err := fmt.Scanln(&userID)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	result, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return
	}
	if rowsAffected == 0 {
		fmt.Println("User ID not found.")
	} else {
		fmt.Println("User deleted successfully.")
	}
}

func changePassword(db *sql.DB) {
	var userID int
	var newPassword string
	fmt.Print("Enter user ID: ")
	_, err := fmt.Scanln(&userID)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Print("Enter new password: ")
	_, err = fmt.Scanln(&newPassword)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	result, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, userID)
	if err != nil {
		fmt.Println("Error updating password:", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		return
	}
	if rowsAffected == 0 {
		fmt.Println("User ID not found.")
	} else {
		fmt.Println("Password updated successfully.")
	}
}
