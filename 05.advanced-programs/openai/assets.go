package main

import (
	"github.com/jessevdk/go-assets"
	"time"
)

var _Assets95155a6c3740bf43201d9858450fce7a49ea750e = "<!DOCTYPE html>\n<html>\n    <head>\n        <title>Input prompt</title>\n        <meta charset=\"utf-8\" />\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n        <style>\n        textarea {\n            width: 80%;\n            height: 50%;\n            resize: none;\n        }\n        </style>\n    </head>\n    <body>\n        <h1>Notes</h1>\n        {{ range .notes }}\n            <div>Note Txt: <pre>{{.Txt}}</pre></div>\n            <div><a href=\"/notedel?id={{.ID}}\">delete</a></div>\n            <br/>\n        {{ end }}\n\n        <h2> Add a new note </h2>\n        <form action=\"/\" method=\"post\">\n            <textarea name=\"note\" id=\"note\" cols=\"80\" rows=\"10\" placeholder=\"Enter your note here\">{{.note}}</textarea>\n            <input type=\"submit\" id=\"submit\" value=\"Submit\" style=\"display:block;\">\n        </form>\n    </body>\n</html>\n\n\n"
var _Assets76bc62f68b0d7d3837ee959f5043ccbee8c22c59 = "<!DOCTYPE html>\n<html>\n    <head>\n        <title>Input prompt</title>\n        <meta charset=\"utf-8\" />\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n        <style>\n        textarea {\n            width: 80%;\n            height: 50%;\n            resize: none;\n        }\n        </style>\n    </head>\n    <body>\n        <h1>Notes</h1>\n        {{ range .notes }}\n            <div>Note Txt: <pre>{{.Txt}}</pre></div>\n            <div><a href=\"/del?id={{.ID}}\">delete</a></div>\n            <br/>\n        {{ end }}\n\n        <h2> Add a new note </h2>\n        <form action=\"/\" method=\"post\">\n            <textarea name=\"note\" id=\"note\" cols=\"80\" rows=\"10\" placeholder=\"Enter your note here\">{{.note}}</textarea>\n            <input type=\"submit\" id=\"submit\" value=\"Submit\" style=\"display:block;\">\n        </form>\n    </body>\n</html>\n\n\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"note.tmpl", "notes.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1679994202, 1679994202160126094),
		Data:     nil,
	}, "/note.tmpl": &assets.File{
		Path:     "/note.tmpl",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1679470090, 1679470090030485656),
		Data:     []byte(_Assets95155a6c3740bf43201d9858450fce7a49ea750e),
	}, "/notes.tmpl": &assets.File{
		Path:     "/notes.tmpl",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1679994202, 1679994202160126094),
		Data:     []byte(_Assets76bc62f68b0d7d3837ee959f5043ccbee8c22c59),
	}}, "")
