package main

import "fmt"

// Declaring constants
const PI float64 = 3.14159265359
const VALUE = 1000

// Multilple Constants Declaration Block
const (
	PRODUCT  = "Ice Cream"
	QUANTITY = 50
	F        = 0.0
)

func main() {
	fmt.Println(PI)
	fmt.Printf("%T\n", PI)
	fmt.Println(VALUE)
	fmt.Printf("%T\n", VALUE)
	fmt.Println(PRODUCT)
	fmt.Printf("%T\n", PRODUCT)
	fmt.Println(QUANTITY)
	fmt.Printf("%T\n", QUANTITY)
	fmt.Println(F)
	fmt.Printf("%T\n", F)
}
