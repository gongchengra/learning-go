package main

// https://gosamples.dev/sqlite-intro/
import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Website struct {
	ID   int64
	Name string
	URL  string
	Rank int64
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS websites(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL,
        rank INTEGER NOT NULL
    );
    `
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(website Website) (*Website, error) {
	res, err := r.db.Exec("INSERT INTO websites(name, url, rank) values(?,?,?)", website.Name, website.URL, website.Rank)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	website.ID = id
	return &website, nil
}

func (r *SQLiteRepository) All() ([]Website, error) {
	rows, err := r.db.Query("SELECT * FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var all []Website
	for rows.Next() {
		var website Website
		if err := rows.Scan(&website.ID, &website.Name, &website.URL, &website.Rank); err != nil {
			return nil, err
		}
		all = append(all, website)
	}
	return all, nil
}

func (r *SQLiteRepository) GetByName(name string) (*Website, error) {
	row := r.db.QueryRow("SELECT * FROM websites WHERE name = ?", name)
	var website Website
	if err := row.Scan(&website.ID, &website.Name, &website.URL, &website.Rank); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &website, nil
}

func (r *SQLiteRepository) Update(id int64, updated Website) (*Website, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE websites SET name = ?, url = ?, rank = ? WHERE id = ?", updated.Name, updated.URL, updated.Rank, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}
	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM websites WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrDeleteFailed
	}
	return err
}

const fileName = "sqlite.db"

func main() {
	os.Remove(fileName)
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	websiteRepository := NewSQLiteRepository(db)
	if err := websiteRepository.Migrate(); err != nil {
		log.Fatal(err)
	}
	gosamples := Website{
		Name: "GOSAMPLES",
		URL:  "https://gosamples.dev",
		Rank: 2,
	}
	golang := Website{
		Name: "Golang official website",
		URL:  "https://golang.org",
		Rank: 1,
	}
	createdGosamples, err := websiteRepository.Create(gosamples)
	if err != nil {
		log.Fatal(err)
	}
	createdGolang, err := websiteRepository.Create(golang)
	if err != nil {
		log.Fatal(err)
	}
	gotGosamples, err := websiteRepository.GetByName("GOSAMPLES")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("get by name: %+v\n", gotGosamples)
	createdGosamples.Rank = 1
	if _, err := websiteRepository.Update(createdGosamples.ID, *createdGosamples); err != nil {
		log.Fatal(err)
	}
	all, err := websiteRepository.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAll websites:\n")
	for _, website := range all {
		fmt.Printf("website: %+v\n", website)
	}
	if err := websiteRepository.Delete(createdGolang.ID); err != nil {
		log.Fatal(err)
	}
	all, err = websiteRepository.All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAll websites:\n")
	for _, website := range all {
		fmt.Printf("website: %+v\n", website)
	}
}
