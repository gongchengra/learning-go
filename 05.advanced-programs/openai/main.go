package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	token := os.Getenv("token")
	if len(token) == 0 {
		fmt.Println("Please provide your openai token.")
		return
	}
	var prompt string
	fmt.Printf("Please enter prompt: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	prompt = scanner.Text()
	c := gogpt.NewClient(token)
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model: gogpt.GPT3TextDavinci003,
		//Model:     gogpt.CodexCodeDavinci002,
		MaxTokens: 500,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
