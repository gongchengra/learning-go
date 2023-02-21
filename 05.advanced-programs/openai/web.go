package main

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
		<head>
				<title>Input Test</title>
		</head>
		<body>
				<form method="post">
						<label for="input">Input Prompt:</label>
						<input type="text" name="input" id="input" />
						<input type="submit" value="Submit" />
				</form>
		</body>
		</html>
		`)
	if r.Method == http.MethodPost {
		//获取用户输入的文本
		input := r.FormValue("input")
		//计算字符数
		token := os.Getenv("token")
		if len(token) == 0 {
			return
		}
		c := gogpt.NewClient(token)
		ctx := context.Background()
		req := gogpt.CompletionRequest{
			Model: gogpt.GPT3TextDavinci003,
			//Model:     gogpt.CodexCodeDavinci002,
			MaxTokens: 500,
			Prompt:    input,
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			return
		}
		fmt.Fprintf(w, resp.Choices[0].Text)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
