package main

import "fmt"

func main() {
	k := 80
	fmt.Println(k, k/9, k%9)
	r, c := k/9, k%9
	r3, c3 := (r/3)*3, (c/3)*3
	fmt.Println(r, c, r3, c3)
	for j := 0; j < 9; j++ {
		cub := r3*9 + c3 + (j/3)*9 + j%3
		fmt.Println(cub)
	}
}
