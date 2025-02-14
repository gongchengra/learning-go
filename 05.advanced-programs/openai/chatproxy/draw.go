package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func drawContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	userID := session.Get(userkey).(int)
	input := c.PostForm("input")

	// Check if content already exists for this input
	existContent, err := checkExist(db, input)
	if err == nil {
		c.HTML(http.StatusOK, "draw.tmpl", gin.H{
			"title":       "Input Prompt",
			"input":       input,
			"imageBase64": existContent.Answer,
		})
		return
	} else {
		log.Println(err)
	}

	// Generate image content via OpenAI
	output, err := draw(input) // expects a base64-encoded image string
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "draw.tmpl", gin.H{"error": "Failed to generate image"})
		return
	}

	// Render the response template
	c.HTML(http.StatusOK, "draw.tmpl", gin.H{
		"title":       "Input Prompt",
		"input":       input,
		"imageBase64": output,
	})

	// Save the content to the database
	err = addContent(db, input, output, userID, 1)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "input.tmpl", gin.H{"error": "Failed to add content"})
		return
	}
}

func draw(input string) (string, error) {
	client := openai.NewClient(option.WithAPIKey(token))
	ctx := context.Background()

	// Set up the image generation parameters.
	params := openai.ImageGenerateParams{
		Prompt:         openai.String(input),
		Model:          openai.F(openai.ImageModelDallE3),
		ResponseFormat: openai.F(openai.ImageGenerateParamsResponseFormatB64JSON),
		N:              openai.Int(1),
	}

	// Generate the image.
	resp, err := client.Images.Generate(ctx, params)
	if err != nil {
		return "", err
	}

	// Return the base64 string for the first generated image.
	return resp.Data[0].B64JSON, nil
}
