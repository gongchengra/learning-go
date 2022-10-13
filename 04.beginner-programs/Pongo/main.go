package main

import (
	"fmt"
	"github.com/flosch/pongo2/v5"
	"log"
)

func main() {
	{
		tpl, err := pongo2.FromString("Hello {{ name }}!")
		if err != nil {
			log.Fatal(err)
		}
		res, err := tpl.Execute(pongo2.Context{"name": "John Doe"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	{
		tpl, err := pongo2.FromString("{{ name }} is a {{ occupation }}")
		if err != nil {
			log.Fatal(err)
		}
		name, occupation := "John Doe", "gardener"
		ctx := pongo2.Context{"name": name, "occupation": occupation}
		res, err := tpl.Execute(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	{
		tpl, err := pongo2.FromFile("message.tpl")
		if err != nil {
			log.Fatal(err)
		}
		name, occupation := "John Doe", "gardener"
		ctx := pongo2.Context{"name": name, "occupation": occupation}
		res, err := tpl.Execute(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	{
		tpl, err := pongo2.FromFile("words.tpl")
		if err != nil {
			log.Fatal(err)
		}
		words := []string{"sky", "blue", "storm", "nice", "barrack", "stone"}
		ctx := pongo2.Context{"words": words}
		res, err := tpl.Execute(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	{
		tpl, err := pongo2.FromFile("wf.tpl")
		if err != nil {
			log.Fatal(err)
		}
		words := []string{"sky", "blue", "storm", "nice", "barrack", "stone"}
		ctx := pongo2.Context{"words": words}
		res, err := tpl.Execute(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	{
		tpl, err := pongo2.FromFile("todos.tpl")
		if err != nil {
			log.Fatal(err)
		}
		todos := []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
			{Title: "Task 4", Done: false},
			{Title: "Task 5", Done: true},
		}
		ctx := pongo2.Context{"todos": todos}
		res, err := tpl.Execute(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}

type Todo struct {
	Title string
	Done  bool
}

type Data struct {
	Todos []Todo
}
