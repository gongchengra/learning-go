package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "chat.db")
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
