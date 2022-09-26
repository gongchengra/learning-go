package main

// https://pkg.go.dev/github.com/yourbasic/bit#hdr-Tutorial
import (
	"fmt"
	"github.com/yourbasic/bit"
	"math"
	"os"
	"strconv"
)

func main() {
	// Sieve of Eratosthenes
	max := 1000
	if len(os.Args) > 1 {
		max, _ = strconv.Atoi(os.Args[1])
	}
	sieve := bit.New().AddRange(2, max)
	sqrtN := int(math.Sqrt(float64(max)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < max; k += p {
			sieve.Delete(k)
		}
	}
	fmt.Println(sieve)
}
