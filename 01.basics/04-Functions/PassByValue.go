package main

import "fmt"

type User struct {
	name       string
	occupation string
}

func inc(x int) {
	x++
	fmt.Printf("inside inc %d\n", x)
}

func change(u User) {
	u.occupation = "driver"
	fmt.Printf("inside change %v\n", u)
}

func incp(x *int) {
	(*x)++
	fmt.Printf("inside inc %d\n", *x)
}

func changep(u *User) {
	u.occupation = "driver"
	fmt.Printf("inside change %v\n", *u)
}

func square(vals []int) {
	for i, val := range vals {
		vals[i] = val * val
	}
}

func squarea(vals [5]int) {
	for i, val := range vals {
		vals[i] = val * val
	}
}

func update(items map[string]int) {
	for k := range items {
		items[k] = 6 * items[k]
	}
	// the same as
	/*
		for k, v := range items {
			items[k] = 6 * v
		}
	*/
	// this one won't work
	/*
		for _,v := range items {
			v = 6 *v
		}
	*/
}

func main() {
	{
		fmt.Println("Pass value")
		x := 10
		fmt.Printf("inside main %d\n", x)
		inc(x)
		fmt.Printf("inside main %d\n", x)
		fmt.Println("---------------------")
		u := User{"John Doe", "gardener"}
		fmt.Printf("inside main %v\n", u)
		change(u)
		fmt.Printf("inside main %v\n", u)
	}
	{
		fmt.Println("\nPass pointer")
		x := 10
		fmt.Printf("inside main %d\n", x)
		incp(&x)
		fmt.Printf("inside main %d\n", x)
		fmt.Println("---------------------")
		u := User{"John Doe", "gardener"}
		fmt.Printf("inside main %v\n", u)
		changep(&u)
		fmt.Printf("inside main %v\n", u)
	}
	{
		vals := []int{1, 2, 3, 4, 5}
		fmt.Printf("%v\n", vals)
		square(vals)
		fmt.Printf("%v\n", vals)
	}
	{
		vals := [5]int{1, 2, 3, 4, 5}
		fmt.Printf("%v\n", vals)
		squarea(vals)
		fmt.Printf("%v\n", vals)
	}
	{
		items := map[string]int{"coins": 1, "pens": 2, "chairs": 4}
		fmt.Printf("%v\n", items)
		update(items)
		fmt.Printf("%v\n", items)
	}
}
