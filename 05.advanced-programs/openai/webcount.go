package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func main() {
	http.HandleFunc("/", submitForm)
	http.ListenAndServe(":8080", nil)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		formStr := `<!DOCTYPE html>
<html>
<head>
<title>input Text Form</title>
</head>
<body>
<h3>Input Text</h3>
<form action="/" method="POST">
<input type="text" name="text" placeholder="input some text here">
<br />
<input type="submit" value="Submit" />
</form>
</body>
</html>`
		fmt.Fprintf(w, formStr)
	} else if r.Method == "POST" {
		r.ParseForm()
		formValue := r.FormValue("text")
		//url-encode
		//         formEncoded := url.QueryEscape(formValue)
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body.Close()
		//计算文字中的字符数
		word := formValue
		wordLen := len(word)
		//计算文字中的中文字符数
		zhCount := 0
		for _, c := range formValue {
			if unicode.Is(unicode.Han, c) {
				zhCount++
			}
		}
		// 计算英文字符数
		enCount := wordLen - zhCount
		// 计算中英文标点数
		punctCount := strings.Count(word, "!") + strings.Count(word, ",") + strings.Count(word, ".") + strings.Count(word, ";") + strings.Count(word, "?")
		//计算空格数
		spaceCount := strings.Count(word, " ")
		//计算段落数
		paraCount := strings.Count(word, "\n\n") + 1
		fmt.Fprintf(w, "您输入的文字中有%d个中文字符、%d个英文字符、%d个段落、%d个空格和%d个中英文标点！ ", zhCount, enCount, paraCount, spaceCount, punctCount)
	}
}
