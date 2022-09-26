package main

// My intuitive thought, without use of channel or gorountine
import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gongchengra/learning-go/02.algorithms/math/SieveOfEratosthenes"
)

func main() {
	max := 1000
	if len(os.Args) > 1 {
		max, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println("Primes less than ", max, ": ")
	start := time.Now()
	var a = make([]bool, max)
	res := []int{}
	for i := 2; i < max; i++ {
		if a[i] == false {
			res = append(res, i)
			for j := i * i; j < max; j = j + i {
				a[j] = true
			}
		}
	}
	fmt.Println(res)
	fmt.Println("seconds it took:", time.Since(start))
	start = time.Now()
	fmt.Println(SieveOfEratosthenes.SieveOfEratosthenes(max))
	fmt.Println("seconds it took:", time.Since(start))
	/*
		for i, j := 2, 0; i < max; i++ {
			if a[i] == false {
				fmt.Printf("%6d ", i)
				j++
				if j%10 == 0 {
					fmt.Println()
				}
			}
		}
	*/
}
