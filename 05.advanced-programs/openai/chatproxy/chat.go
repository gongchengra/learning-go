package main

import (
	"log"

	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
)

func chat(input string, assist string) (output string) {
	c := gpt35.NewClient(token)
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleAssistant,
				Content: assist,
			},
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
