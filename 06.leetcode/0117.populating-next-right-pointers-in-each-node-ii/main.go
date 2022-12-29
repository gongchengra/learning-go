package main

import "fmt"

func main() {
	n1 := &Node{}
	n1.Val = 1
	n2 := &Node{}
	n2.Val = 2
	n3 := &Node{}
	n3.Val = 3
	n4 := &Node{}
	n4.Val = 4
	n5 := &Node{}
	n5.Val = 5
	n7 := &Node{}
	n7.Val = 7
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	n3.Right = n7
	connect(n1)
	fmt.Println(n1.Next)
	fmt.Println(n2.Next.Val, n3.Next)
	fmt.Println(n4.Next.Val, n5.Next.Val, n7.Next)
}
