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
		os.Exit(1) // Exit if environment variable doesn't exist
	}
	return value
}

const (
	apiBaseURLTemplate = "https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-int8"
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

func chat(input, assist string) (output string) {
	accountID := getEnvVariable("ACCOUNT_ID")
	apiToken := getEnvVariable("API_TOKEN")

	url := fmt.Sprintf(apiBaseURLTemplate, accountID)
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
		fmt.Println(err)
		return
	}
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(respBody))
        return
	}

	var chatResp ChatResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		fmt.Errorf("error unmarshaling response: %s, response body: %s", err, string(respBody))
        return
	}

	if chatResp.Success && len(chatResp.Result.Response) > 0 {
		output = chatResp.Result.Response
	} else {
		fmt.Errorf("empty response body or failure in response: %s", string(respBody))
        return
	}

	return output
}

