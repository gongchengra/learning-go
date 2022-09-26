package main

import "fmt"

type output func(string) string

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	var f output
	f = hello
	fmt.Println(f("Peter"))
}
