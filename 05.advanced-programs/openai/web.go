package main

import (
	"fmt"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
    <head>
        <title>Input prompt</title>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
    </head>
    <body>
		<h1>Chat with AI</h1>
		<p> 在下面文本框中输入指令，例如：“把下面一段话翻译成英文： 在下面文本框中输入指令。” 或者“用英文写一封邮件，回复Anderson先生，解释由于快递原因，他的床垫需要下个月底才到送到他家。我们表示非常抱歉。”，关于如何制作清晰有效的Prompt指南，请参考<a href="https://openai.wiki/chatgpt-prompting-guide-book.html" target="_blank">https://openai.wiki/chatgpt-prompting-guide-book.html</a> 网站。
        <form method="post">
            <textarea name="input" id="input" cols="80" rows="10" placeholder="Enter your prompt here"></textarea>
            <input type="submit" id="submit" value="Submit" style="display:block;">
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
		c := gpt35.NewClient(token)
		req := &gpt35.Request{
			Model: gpt35.ModelGpt35Turbo,
			Messages: []*gpt35.Message{
				{
					Role:    gpt35.RoleUser,
					Content: input,
				},
			},
		}
		resp, err := c.GetChat(req)
		if err != nil {
			panic(err)
		}
		content := resp.Choices[0].Message.Content
		fmt.Fprintf(w, `<br><pre>%s</pre>`, input)
		fmt.Fprintf(w, `<br><pre>%s</pre>`, content)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
