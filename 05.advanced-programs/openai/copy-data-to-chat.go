package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open data.db and chat.db
	dataDB, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println("Error while opening data.db")
		return
	}
	defer dataDB.Close()
	chatDB, err := sql.Open("sqlite3", "chat.db")
	if err != nil {
		fmt.Println("Error while opening chat.db")
		return
	}
	defer chatDB.Close()
	// Query records from contents table in data.db
	rows, err := dataDB.Query("SELECT prompt, answer, userid FROM contents")
	if err != nil {
		fmt.Println("Error while querying contents table in data.db")
		return
	}
	defer rows.Close()
	// Loop through each row and insert into chat.db if prompt does not exist
	for rows.Next() {
		var prompt, answer string
		var userid int
		err = rows.Scan(&prompt, &answer, &userid)
		if err != nil {
			fmt.Println("Error while scanning row in contents table")
			return
		}
		// Check if prompt already exists in chat.db
		var exists bool
		err = chatDB.QueryRow("SELECT EXISTS(SELECT 1 FROM contents WHERE prompt=?)", prompt).Scan(&exists)
		if err != nil {
			fmt.Println("Error while checking if prompt exists in contents table in chat.db")
			return
		}
		if !exists {
			// Insert record into contents table in chat.db
			_, err = chatDB.Exec("INSERT INTO contents (prompt, answer, userid) VALUES (?, ?, ?)", prompt, answer, userid)
			if err != nil {
				fmt.Println("Error while inserting record into contents table in chat.db")
				return
			}
		}
	}
	fmt.Println("Records copied successfully!")
}
