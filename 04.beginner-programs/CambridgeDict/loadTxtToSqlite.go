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
    CREATE TABLE IF NOT EXISTS words(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        word TEXT NOT NULL UNIQUE,
        definition TEXT NOT NULL
    );
    `
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(word string, def string) (id string, err error) {
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

func (r *SQLiteRepository) GetByWord(word string) (def string, err error) {
	row := r.db.QueryRow("SELECT definition FROM words WHERE word = ?", word)
	if err := row.Scan(&def); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotExists
		}
		return "", err
	}
	return def, nil
}

/*
const fileName = "words.db"

func main() {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	websiteRepository := NewSQLiteRepository(db)
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		filename := info.Name()
		ext := filepath.Ext(filename)
		if info.IsDir() || ext != ".txt" {
			return nil
		}
		word := strings.TrimSuffix(filename, filepath.Ext(filename))
		definition, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		//         fmt.Println(word, string(definition))
		id, err := websiteRepository.Create(word, string(definition))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return nil
		}
		fmt.Println(id + "created")
		return nil
	})
	//     if err := websiteRepository.Migrate(); err != nil {
	//         log.Fatal(err)
	//     }
	//     word := "zoo"
	//     filename := word + ".txt"
	//     content, err := os.ReadFile(filename)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     _, err = websiteRepository.Create(word, string(content))
	//     gotdef, err := websiteRepository.GetByWord(word)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     fmt.Printf("get by word: %+v\n", gotdef)
}
*/
