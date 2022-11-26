package main

import "fmt"

func main() {
	// Declaring variables
	var i int
	var s string
	// Initializing Variables
	i = 20
	s = "Some String"
	fmt.Println(i)
	fmt.Println(s)
	// Creating and initializing variables
	var k int = 35
	fmt.Println(k)
	//Short variable declaration
	j := 50
	fmt.Println(j)
	// Declaring multiple variables
	firstName, lastName := "FirstName", "LastName"
	fmt.Println(firstName + lastName)
	// Variable Declaration Block
	var (
		name = "Donald Duck"
		age  = 50
	)
	fmt.Println(name)
	fmt.Println(age)
	{
		a := [3]int{1, 2, 3}
		for k, v := range a {
			if k == 0 {
				a[0], a[1] = 100, 200
			}
			a[k] = 100 + v
		}
		fmt.Print(a)
	}
	{
		a := []int{1, 2, 3}
		//         a[0], a[1] = 100, 200
		for k, v := range a {
			if k == 0 {
				a[0], a[1] = 100, 200
				//                 a[k], a[k+1] = 100, 200
			}
			a[k] = 100 + v
		}
		fmt.Print(a) //数组    101  102  103
	}
}
