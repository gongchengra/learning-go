package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ts := httptest.NewServer(writeHTML(`
<head>
    <title>Home page</title>
</head>
<body>
    <p>Hello there!</a>
</body>
    `))
	defer ts.Close()
	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Title(&title),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)
}
