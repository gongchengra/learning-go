package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Print("Enter database path: ")
	var dbPath string
	fmt.Scanln(&dbPath)
	if dbPath == "" {
		dbPath = "chat.db"
	}
	setConsecutiveId(dbPath)
	setMaxId(dbPath)
}

func setConsecutiveId(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare("SELECT id FROM contents ORDER BY id")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id, newId int = 0, 1
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		_, err = tx.Exec("UPDATE contents SET id = ? WHERE id = ?", newId, id)
		if err != nil {
			panic(err)
		}
		newId++
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Id consecutive number updated successfully!")
}

func setMaxId(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare("SELECT MAX(id) FROM contents")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var maxId int // added new variable to store max id
	for rows.Next() {
		err = rows.Scan(&maxId)
		if err != nil {
			panic(err)
		}
	}
	_, err = tx.Exec("UPDATE sqlite_sequence SET seq = ? WHERE name = 'contents'", maxId)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	fmt.Println("Maximum id reset successfully!")
}
