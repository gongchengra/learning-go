package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func getEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		fmt.Printf("Error: %s environment variable not set\n", key)
	}
	return value
}

const (
	apiBaseURLTemplate = "https://api.cloudflare.com/client/v4/accounts/%s/ai/run/"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Result struct {
		Response string `json:"response"`
	} `json:"result"`
	Success  bool        `json:"success"`
	Errors   []struct{}  `json:"errors"`
	Messages []struct{}  `json:"messages"`
}

func chat(model, input, assist string) (output string, err error) {
	accountID := getEnvVariable("ACCOUNT_ID")
	apiToken := getEnvVariable("API_TOKEN")

	url := fmt.Sprintf(apiBaseURLTemplate, accountID) + model
	messages := []ChatMessage{
		{
			Role:    "system",
			Content: assist,
		},
		{
			Role:    "user",
			Content: input,
		},
	}
	requestBody, err := json.Marshal(ChatRequest{Messages: messages})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	var chatResp ChatResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %s, response body: %s", err, string(respBody))
	}

	if chatResp.Success && len(chatResp.Result.Response) > 0 {
		output = chatResp.Result.Response
	} else {
		return "", fmt.Errorf("empty response body or failure in response: %s", string(respBody))
	}

	return output, nil
}

func main() {
	model := "@cf/meta/llama-2-7b-chat-int8"
	input := "Write a short story about a llama that goes on a journey to find an orange cloud."
	assist := "You are a friendly assistant that helps write stories."

	output, err := chat(model, input, assist)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", output)
}
