package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

const (
	// cse
	cx = "25100eeabd01c4e0b"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	apikey := os.Getenv("API_KEY")
	client := &http.Client{Transport: &transport.APIKey{Key: apikey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range fileLines {
		query := "小说 " + line
		resp, err := svc.Cse.List().Cx(cx).Q(query).Do()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line, resp.SearchInformation.TotalResults)
	}
}
