package main

import (
	"context"
	"fmt"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func chat(input string, assist string) (output string) {
	client := openai.NewClient(option.WithAPIKey(token))

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.AssistantMessage(assist),
		openai.UserMessage(input),
	}

	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(openai.ChatModelO3Mini),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	output = chatCompletion.Choices[0].Message.Content
	return output
}
