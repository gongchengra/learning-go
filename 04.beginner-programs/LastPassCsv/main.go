package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

// Compile templates on start of the application
var index = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Lastpass Search</title>
  </head>
  <body>
    <form enctype="multipart/form-data" action="/" method="post">
	  <label for="search">Url:</label>
      <input type="text" name="search" />
      <input type="submit" value="search" />
    </form>
    <form enctype="multipart/form-data" action="/addRecord" method="post">
	  <label for="url">Url:</label>
      <input type="url" name="url" />
	  <label for="username">Username:</label>
      <input type="text" name="username" />
	  <label for="password">Password:</label>
      <input type="password" name="password" />
      <input type="submit" value="add" />
    </form>
	<p>{{.}}</p>
  </body>
</html>
`))
var list = template.Must(
	template.New("list.html").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>List File</title>
  </head>
  <body>
      <h1>Results</h1>
      {{range $rec := .}}
      <div>
	  <p><a href="{{$rec.Url}}" target="_blank">{{$rec.Url}}</a></p>
	  <p>{{$rec.Username}}</p>
	  <p>{{$rec.Password}}</p>
	  </div>
      {{end}}
  </body>
</html>
`),
)

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

type Record struct {
	Url      string
	Username string
	Password string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		index.ExecuteTemplate(w, "index.html", nil)
	case "POST":
		search := r.FormValue("search")
		records, _ := readData("lastpass.log")
		res := []Record{}
		for _, record := range records {
			if strings.Contains(record[0], search) {
				rec := Record{
					Url:      record[0],
					Username: record[1],
					Password: record[2],
				}
				res = append(res, rec)
			}
		}
		list.ExecuteTemplate(w, "list.html", res)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	username := r.FormValue("username")
	password := r.FormValue("password")
	records, _ := readData("lastpass.log")
	res := []string{url, username, password, "", "", "", "", ""}
	records = append(records, res)
	f, err := os.Create("lastpass.log")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	write := csv.NewWriter(f)
	err = write.WriteAll(records) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}
	res = append(res, " added")
	index.ExecuteTemplate(w, "index.html", res)
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func main() {
	app := new(application)
	app.auth.username = os.Getenv("AUTH_USERNAME")
	app.auth.password = os.Getenv("AUTH_PASSWORD")
	if app.auth.username == "" {
		log.Fatal("basic auth username must be provided")
	}
	if app.auth.password == "" {
		log.Fatal("basic auth password must be provided")
	}
	// Index route
	http.HandleFunc("/", app.basicAuth(indexHandler))
	http.HandleFunc("/addRecord", app.basicAuth(addHandler))
	http.ListenAndServe(":8080", nil)
}
