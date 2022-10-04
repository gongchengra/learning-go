package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed data/words.txt
	data []byte
)

func main() {
	fmt.Println(string(data))
	fmt.Println("----------------------")
	{
		words := strings.Split(string(data), "\n")
		for _, w := range words {
			fmt.Println(string(w))
		}
	}
	fmt.Println("----------------------")
	{
		words := bytes.Split(data, []byte{'\n'})
		for _, w := range words {
			fmt.Println(string(w))
		}
	}
}
