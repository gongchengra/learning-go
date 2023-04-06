package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite3 database
	db, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()
	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error beginning transaction:", err)
		return
	}
	// Query the contents table
	rows, err := tx.Query("SELECT id, prompt FROM contents")
	if err != nil {
		fmt.Println("Error querying contents table:", err)
		tx.Rollback()
		return
	}
	// Create a map to keep track of prompts
	prompts := make(map[string]int)
	// Iterate through the rows
	for rows.Next() {
		var id int
		var prompt string
		if err := rows.Scan(&id, &prompt); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		// Check if prompt already exists in map
		if _, ok := prompts[prompt]; ok {
			// Prompt already exists, mark current row as deleted
			if _, err := tx.Exec("DELETE FROM contents WHERE id = ?", id); err != nil {
				fmt.Println("Error marking row as deleted:", err)
				tx.Rollback()
				return
			}
			fmt.Println("Deleted duplicate row with prompt:", prompt)
		} else {
			// Prompt doesn't exist, add to map
			prompts[prompt] = id
		}
	}
	// Commit changes
	if err := tx.Commit(); err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}
}
