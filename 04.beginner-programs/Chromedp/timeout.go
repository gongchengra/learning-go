package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	url := "http://webcode.me"
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text("body", &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(res))
}
