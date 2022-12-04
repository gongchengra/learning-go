package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle(
		[][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'}}))
}
