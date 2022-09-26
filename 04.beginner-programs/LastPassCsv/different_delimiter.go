package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("users1.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(records)
}
