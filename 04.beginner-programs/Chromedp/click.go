package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"log"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	url := "http://webcode.me/click.html"
	var ua string
	err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPhone11),
		chromedp.Navigate(url),
		chromedp.Click("button", chromedp.NodeVisible),
		chromedp.Text("#output", &ua),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User agent: %s\n", ua)
}
