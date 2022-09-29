package main

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate = errors.New("record already exists")
	ErrNotExists = errors.New("row not exists")
)

type sqlitedb struct {
	db *sql.DB
}

func Newsqlitedb(db *sql.DB) *sqlitedb {
	return &sqlitedb{
		db: db,
	}
}

func (r *sqlitedb) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS words(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        word TEXT NOT NULL UNIQUE,
        definition TEXT NOT NULL
    );
    `
	_, err := r.db.Exec(query)
	return err
}

func (r *sqlitedb) Create(word string, def string) (id string, err error) {
	res, err := r.db.Exec("INSERT INTO words(word, definition) values(?,?)", word, def)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return "", ErrDuplicate
			}
		}
		return "", err
	}
	insert, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(insert)), nil
}

func (r *sqlitedb) GetByWord(word string) (def string, err error) {
	row := r.db.QueryRow("SELECT definition FROM words WHERE word = ?", word)
	if err := row.Scan(&def); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotExists
		}
		return "", err
	}
	return def, nil
}
