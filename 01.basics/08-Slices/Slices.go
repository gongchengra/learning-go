package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Slices are dynamic datastructures that can grow and shrink as you see fit.

func uniq(vals []int) []int {
	uvals := []int{}
	seen := make(map[int]bool)
	for _, val := range vals {
		if _, in := seen[val]; !in {
			seen[val] = true
			uvals = append(uvals, val)
		}
	}
	return uvals
}

func main() {
	p1 := false
	//     p1 := false
	//     p2 := true
	p2 := false
	//     p3 := true
	p3 := false
	p4 := true
	if p1 {
		// Create an empty slice
		var x []int
		fmt.Println(reflect.ValueOf(x).Kind())
		// Creating a slice using the make function
		var y = make([]string, 10, 20)
		fmt.Printf("y \tLen: %v \tCap: %v\n", len(y), cap(y))
		// Initialize the slice with values using a slice literal
		var z = []int{10, 20, 30, 40}
		fmt.Printf("z \tLen: %v \tCap: %v\n", len(z), cap(z))
		fmt.Println(z)
		// Creating a Slice using the new keyword
		var a = new([50]int)[0:10]
		fmt.Printf("a \tLen: %v \tCap: %v\n", len(a), cap(a))
		fmt.Println(a)
		// Add items using the append function
		var b = make([]int, 1, 10)
		fmt.Println(b)
		b = append(b, 20)
		fmt.Println(b)
		// Access slice items
		var c = []int{10, 20, 30, 40}
		fmt.Println(c[0])
		fmt.Println(c[0:3])
		// Change item values
		var d = []int{10, 20, 30, 40}
		fmt.Println(d)
		d[1] = 35
		fmt.Println(d)
		// Copy slice into another slice
		var e = []int{10, 20, 30, 40}
		var f = []int{50, 60, 70, 80}
		copy(e, f)
		fmt.Println("E: ", e)
		// Append a slice to an existing one
		var g = []int{10, 20, 30, 40}
		var h = []int{50, 60, 70, 80}
		g = append(g, h...)
		fmt.Println(g)
	}
	if p2 {
		arr := make([]int, 0)
		for i, c := 0, 0; i < 200000; i++ {
			if cap(arr) != c {
				fmt.Println(i, c, cap(arr), float32(cap(arr))/float32(c))
				c = cap(arr)
			}
			//             fmt.Println("len 为 ", len(arr), "cap 为 ", cap(arr))
			arr = append(arr, i)
		}
	}
	if p3 {
		// w is slice of slices
		w := make([][]string, 3)
		w1 := make([]string, 4)
		w1[0] = "war"
		w1[1] = "water"
		w1[2] = "wrath"
		w1[3] = "wrong"
		w2 := make([]string, 3)
		w2[0] = "car"
		w2[1] = "cup"
		w2[2] = "cloud"
		w3 := make([]string, 2)
		w3[0] = "boy"
		w3[1] = "brown"
		w[0] = w1
		w[1] = w2
		w[2] = w3
		fmt.Println(w)
	}
	// https://zetcode.com/golang/slice/
	if p4 {
		vals := []int{1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 8, 8, 9, 9}
		uvals := uniq(vals)
		fmt.Printf("Original slice: %v\n", vals)
		fmt.Printf("Unique slice: %v\n", uvals)
	}
	{
		words := []string{"falcon", "bold", "bear", "sky", "cloud", "ocean"}
		vals := []int{4, 2, 1, 5, 6, 8, 0, -3}
		sort.Strings(words)
		sort.Ints(vals)
		fmt.Println(words)
		fmt.Println(vals)
	}
	{
		vals := []int{1, 2, 3, 4, 5, 6}
		vals2 := vals
		vals2[0] = 11
		vals2[1] = 22
		fmt.Println(vals)
		fmt.Println(vals2)
	}
}
