package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := "http://webcode.me"
	var data string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
