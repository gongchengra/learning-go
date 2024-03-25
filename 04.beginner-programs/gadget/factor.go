package main

import (
	"fmt"
	"math/big"
)

// primeFactors finds prime factors of a given big.Int number.

func primeFactors(number *big.Int) []*big.Int {
	// A slice to store the prime factors
	var factors []*big.Int
	// The variable 'i' is our potential factor, starting at 2.
	i := big.NewInt(2)
	// Variables to hold intermediate results
	var divMod big.Int
	var zero big.Int
	// Continue dividing 'number' by 'i' until 'i' squared is greater than 'number'.
	for iSquared := new(big.Int).Mul(i, i); iSquared.Cmp(number) <= 0; iSquared.Mul(i, i) {
		// While 'i' divides 'number', append 'i' to the list and divide 'number' by 'i'.
		for divMod.Mod(number, i).Cmp(&zero) == 0 {
			factors = append(factors, new(big.Int).Set(i))
			number.Div(number, i)
		}
		i.Add(i, big.NewInt(1))
	}
	// If the remaining 'number' is not 1, it is a prime factor itself.
	if number.Cmp(big.NewInt(1)) != 0 {
		factors = append(factors, number)
	}
	return factors
}

func main() {
	// The input number.
	number := new(big.Int)
	//     number.SetString("50296446669475900103173373917676050797315238472655018625792237958753851302131", 10)
	number.SetString("50296446669475902694584875652761273678224636096879696755629890223059048071168", 10)
	// Find the prime factors.
	factors := primeFactors(number)
	// Print the prime factors.
	fmt.Println("Prime factors:")
	for _, factor := range factors {
		fmt.Println(factor)
	}
}
