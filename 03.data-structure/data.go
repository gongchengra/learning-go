package main

import (
	"fmt"
	"laogongshuo.com/data/LinkedList"
)

func main() {
	l := LinkedList.LinkedList{}
	l.Append(5)
	l.Append(50)
	l.Append(501)
	l.Prepend(404)
	l.Prepend(40)
	l.Prepend(4)
	l.Display()
	l.RemoveAtBeg()
	l.RemoveAtEnd()
	l.Display()
	fmt.Println(l.Count())
	l.Reverse()
	l.Display()
	l.DisplayReverse()
}
