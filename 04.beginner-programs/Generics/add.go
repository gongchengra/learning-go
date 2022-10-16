package main

import "fmt"

func Add[T int | float32 | float64 | string](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add[int](1, 2))
	fmt.Println(Add[float32](1.0, 2.0))
	fmt.Println(Add[string]("Hello ", "world!"))
}
