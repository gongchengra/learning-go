package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	url := "http://webcode.me"
	var buf []byte
	if err := chromedp.Run(ctx, ElementScreenshot(url, "body", &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("body.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	if err := chromedp.Run(ctx, FullScreenshot(url, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("full.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("screenshots created")
}

func ElementScreenshot(url, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

func FullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(res, quality),
	}
}
