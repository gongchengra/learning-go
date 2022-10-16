package main

import "fmt"

type Number interface {
	~int32 | int64 | float64
	// this one won't work:  int32 | int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type testInt int32

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	myints := map[string]testInt{
		"first":  34,
		"second": 12,
	}
	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("Generic Sums with Constraint: %v , %v and %v\n",
		SumNumbers(ints),
		SumNumbers(myints),
		SumNumbers(floats))
}
