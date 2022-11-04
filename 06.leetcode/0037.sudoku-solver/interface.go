package main

import "fmt"

func main() {
	kv := make([]interface{}, 2)
	kv[0] = 55
	kv[1] = "1234"
	fmt.Println(kv[0].(int), kv[1].(string))
}
