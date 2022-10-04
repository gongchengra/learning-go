package main

import (
	"embed"
	"fmt"
)

//go:embed data/*
var f embed.FS

func main() {
	langs, _ := f.ReadFile("data/langs.txt")
	fmt.Println(string(langs))
	words, _ := f.ReadFile("data/words.txt")
	fmt.Println(string(words))
}
