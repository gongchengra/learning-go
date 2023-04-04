package main

import (
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"log"
)

func chat(input string) (output string) {
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
		log.Println(err)
		return
	}
	output = resp.Choices[0].Message.Content
	return output
}
