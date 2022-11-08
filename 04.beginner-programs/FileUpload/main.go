package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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
    <title>Upload and Download Files</title>
  </head>
  <body>
      <div>
          <a href="/upload" target="_blank">Upload file</a>
      </div>
      <div>
          <a href="/list" target="_blank">Browse and Download files</a>
      </div>
  </body>
</html>
`))
var upload = template.Must(template.New("upload.html").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Upload File</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
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
      <h1>Uploaded Files</h1>
      {{range $key, $val := .}}
      <div><a href = "/files/{{$val.href}}" target="_blank">{{$val.name}}</a></div>
      {{end}}
  </body>
</html>
`),
)

const uploadPath = "./data"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 100 MB files
	r.ParseMultipartForm(100 << 20)
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	// Create file
	dst, err := os.Create("./data/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		upload.ExecuteTemplate(w, "upload.html", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func SortFileNameAscend(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(uploadPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	if dirs, err := f.Readdir(-1); err == nil {
		SortFileNameAscend(dirs)
		files := make([]map[string]string, len(dirs)+1)
		for i, d := range dirs {
			href := d.Name()
			if d.IsDir() {
				href += "/"
			}
			files[i+1] = map[string]string{
				"name": d.Name(),
				"href": href,
			}
		}
		list.ExecuteTemplate(w, "list.html", files)
	}
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.ExecuteTemplate(w, "index.html", nil)
	})
	// Upload route
	http.HandleFunc("/upload", app.basicAuth(uploadHandler))
	// Download route
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	// List route
	http.HandleFunc("/list", app.basicAuth(listHandler))
	http.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", `attachment;`)
		http.ServeFile(w, r, filepath.Join(".", r.URL.Path))
	})
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", `attachment;`)
		http.ServeFile(w, r, filepath.Join(".", strings.Replace(r.URL.Path, "download", "data", 1)))
	})

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
