package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

func drawContentHandler(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	session := sessions.Default(c)
	userID := session.Get(userkey).(int)
	input := c.PostForm("input")
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
	output, err := draw(input) // Now expects a base64 image string
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "draw.tmpl", gin.H{"error": "Failed to generate image"})
		return
	}
	c.HTML(http.StatusOK, "draw.tmpl", gin.H{
		"title":       "Input Prompt",
		"input":       input,
		"imageBase64": output,
	})
	err = addContent(db, input, output, userID, 1)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "input.tmpl", gin.H{"error": "Failed to add content"})
		return
	}
}

func draw(input string) (string, error) {
	c := openai.NewClient(token)
	ctx := context.Background()

	reqBase64 := openai.ImageRequest{
		Model:          openai.CreateImageModelDallE3,
		N:              1,
		Quality:        openai.CreateImageQualityHD,
		Size:           openai.CreateImageSize1024x1024,
		Style:          openai.CreateImageStyleVivid,
		Prompt:         input,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		return "", err
	}
	return respBase64.Data[0].B64JSON, nil
}
