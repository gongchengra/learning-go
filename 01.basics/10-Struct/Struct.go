package main

import "fmt"

// Declaring a Struct

type Animal struct {
	name   string
	weight int
}

func newAnimal(n string, w int) *Animal {
	return &Animal{name: n, weight: w}
}

func main() {
	// Creating an instance of a struct
	var dog Animal
	dog.name = "Dog"
	dog.weight = 40
	fmt.Println(dog)
	// Creating an instance using struct literate
	var cat = Animal{name: "Cat", weight: 5}
	fmt.Println(cat)
	// Creating an instance using the new keyword
	var bird = new(Animal)
	bird.name = "Bird"
	bird.weight = 1
	fmt.Println(bird)
	// Creating an instance using the pointer address operator
	var monkey = &Animal{name: "Monkey", weight: 10}
	fmt.Println(monkey)
	// Comparing struct instances
	fmt.Println(bird == monkey)
	// Copying struct type using pointer reference
	monkey2 := monkey
	monkey2.name = "Monkey2"
	fmt.Println(monkey2)
	p1 := newAnimal("Pig", 50)
	fmt.Println(p1, *p1)
	p2 := newAnimal("Pig", 50)
	fmt.Println(p2, *p2)
	fmt.Println(p1 == p2, *p1 == *p2) //false true
	p3 := Animal{name: "Pig", weight: 30}
	p4 := Animal{name: "Pig", weight: 30}
	fmt.Println(p3 == p4) //true
}
