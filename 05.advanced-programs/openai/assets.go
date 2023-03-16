package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\n<html>\n    <head>\n        <title>Input prompt</title>\n        <meta charset=\"utf-8\" />\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    </head>\n    <body>\n\t\t<h1>Chat with AI</h1>\n        <p> Enter the command in the text box below, such as \"Translate the following paragraph into Chinese: Enter the command in the text box below.\" or \"Write an email in French to Mr. Anderson, explaining that due to delivery reasons, his mattress will only be delivered to his home at the end of next month. We apologize for the inconvenience.\" For guidance on how to create clear and effective prompt guidelines, please refer to the website <a href=\"https://openai.wiki/chatgpt-prompting-guide-book.html\" target=\"_blank\">https://openai.wiki/chatgpt-prompting-guide-book.html</a>.\n        <form method=\"post\">\n            <textarea name=\"input\" id=\"input\" cols=\"80\" rows=\"10\" placeholder=\"Enter your prompt here\">{{.input}}</textarea>\n            <input type=\"submit\" id=\"submit\" value=\"Submit\" style=\"display:block;\">\n        </form>\n        {{.output}}\n    </body>\n</html>\n\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1678954897, 1678954897529129360),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1678968169, 1678968169796685061),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1678969000, 1678969000890528832),
		Data:     nil,
	}}, "")
