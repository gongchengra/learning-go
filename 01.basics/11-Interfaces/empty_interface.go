package main

// https://zetcode.com/golang/interface/
import (
	"fmt"
	"log"
)

type Body struct{ Msg interface{} }

func main() {
	b := Body{"Hello there"}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
	b.Msg = 5
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
	b.Msg = []int{5, 6, 7}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
	b.Msg = 5.1
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
	user := make(map[string]interface{}, 0)
	user["name"] = "John Doe"
	user["age"] = 21
	user["weight"] = 70.3
	age, ok := user["age"].(int)
	if !ok {
		log.Fatal("assert failed")
	}
	user["age"] = age + 1
	fmt.Printf("%+v", user)
}
