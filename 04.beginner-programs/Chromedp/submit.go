package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 6*time.Second)
	defer cancel()
	url := "http://webcode.me/submit/"
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.SendKeys("input[name=name]", "Lucia"),
		chromedp.SendKeys("input[name=message]", "Hello!"),
		// chromedp.Click("button", chromedp.NodeVisible),
		chromedp.Submit("input[name=name]"),
		chromedp.Text("*", &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
