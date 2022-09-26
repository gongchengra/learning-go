package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	e3 := l.InsertBefore(3, e4)
	e2 := l.InsertAfter(2, e1)
	//	fmt.Println(l.Front(), e1, e2)
	fmt.Print(l.Front().Value, e1.Value, e2.Value)
	fmt.Print(e3.Value, e4.Value, l.Back().Value)
	fmt.Println(l.Len())
	fmt.Println("Move and Print")
	l.MoveAfter(e3, e4)
	l.MoveBefore(e2, e1)
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("Move and Print")
	l.MoveToBack(e4)
	l.MoveToFront(e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	for i := 0; i < 20; i++ {
		l.PushBack(i)
	}
	fmt.Println(l.Back().Value, e4.Value)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}
