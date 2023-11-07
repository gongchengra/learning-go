package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func chat(input string, assist string) (output string) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4TurboPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: assist,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	output = resp.Choices[0].Message.Content
	return output
}
