package main

import (
	"bufio"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("token")
	if len(token) == 0 {
		fmt.Println("Please provide your openai token.")
		return
	}
	var prompt string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter prompt: ")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "EOF" {
			break
		}
		prompt += line
	}
	fmt.Println("waiting: ")
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
