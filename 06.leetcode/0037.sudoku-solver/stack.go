package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	stack.PushBack(11)
	stack.PushBack("abc")
	fmt.Println(stack.Back().Value.(string))
	stack.Remove(stack.Back())
	fmt.Println(stack.Back().Value.(int))
}
