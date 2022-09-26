package main

// https://zetcode.com/golang/interface/

type Animal interface{ Sound() string }

type Dog struct{}

func (d Dog) Sound() string { return "Woof!" }

type Cat struct{}

func (c Cat) Sound() string { return "Meow!" }

type Cow struct{}

func (l Cow) Sound() string { return "Moo!" }

func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		println(animal.Sound())
	}
}
