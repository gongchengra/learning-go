package main

import (
	"fmt"
	"net/http"
	"os"

	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/gin-gonic/gin"
)

func chat(input string) (output string) {
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
		fmt.Println(err)
		panic(err)
	}
	output = resp.Choices[0].Message.Content
	return output
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.LoadHTMLGlob("html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"input": "",
		})
	})
	r.POST("/", func(c *gin.Context) {
		input := c.PostForm("input")
		output := chat(input)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"input":  input,
			"output": output,
		})
	})
	r.Run(":8081")
}
