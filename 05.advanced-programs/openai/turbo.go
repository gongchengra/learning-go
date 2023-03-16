package main

import (
	"bufio"
	"fmt"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	token := os.Getenv("token")
	if len(token) == 0 {
		fmt.Println("Please provide your openai token.")
		return
	}
	var prompt string
	fmt.Println("Please enter prompt: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "EOF" {
			break
		}
		prompt += line
	}
	fmt.Println("waiting: ")
	recLen := 0
	if len(os.Args) > 1 {
		recLen, _ = strconv.Atoi(os.Args[1])
	}
	assist := ""
	if recLen > 0 {
		data, err := os.ReadFile(strings.TrimSpace(getLastLineWithSeek("my.log")))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
		assist = string(data)
	}
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
				Content: prompt,
			},
		},
	}
	resp, err := c.GetChat(req)
	if err != nil {
		panic(err)
	}
	content := resp.Choices[0].Message.Content
	t := time.Now().Unix()
	name := fmt.Sprintf("%d.log", t)
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d.log\n", t))
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(content)
	println(content)
	println(resp.Usage.PromptTokens)
	println(resp.Usage.CompletionTokens)
	println(resp.Usage.TotalTokens)
}

func getLastLineWithSeek(filepath string) string {
	fileHandle, err := os.Open(filepath)
	if err != nil {
		panic("Cannot open file")
		os.Exit(1)
	}
	defer fileHandle.Close()
	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)
		char := make([]byte, 1)
		fileHandle.Read(char)
		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}
		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way
		if cursor == -filesize {                       // stop if we are at the begining
			break
		}
	}
	return line
}
