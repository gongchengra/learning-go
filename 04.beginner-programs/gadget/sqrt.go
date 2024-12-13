package main

import (
    "fmt"
    "math/big"
)

func calculateSqrt(x float64, precision int) *big.Float {
    // Set the precision for big.Float
    bigFloatVal := new(big.Float).SetPrec(uint(precision))

    // Set the value to big.Float
    bigFloatVal.SetFloat64(x)

    // Calculate the square root
    sqrtVal := new(big.Float).Sqrt(bigFloatVal)

    return sqrtVal
}

func main() {
    var number float64
    var precision int

    // Get user input
    fmt.Print("Enter a number: ")
    fmt.Scan(&number)

    fmt.Print("Enter the desired precision (in bits): ")
    fmt.Scan(&precision)

    sqrtResult := calculateSqrt(number, precision)

    // Print the result
    fmt.Printf("The square root of %.2f with %d bits of precision is: %s\n", number, precision, sqrtResult.Text('f', precision))
}

