package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{[]byte{'a', 'b', 'c', 'e'}, []byte{'s', 'f', 'c', 's'}, []byte{'a', 'd', 'e', 'f'}}, "abcced"))
	fmt.Println(exist([][]byte{[]byte{'a', 'b', 'c', 'e'}, []byte{'s', 'f', 'c', 's'}, []byte{'a', 'd', 'e', 'e'}}, "see"))
}
