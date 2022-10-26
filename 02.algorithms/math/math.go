package main

import (
	"fmt"
	"math"

	"github.com/gongchengra/learning-go/02.algorithms/math/EuclideanAlgorithm"
	"github.com/gongchengra/learning-go/02.algorithms/math/Factorial"
	"github.com/gongchengra/learning-go/02.algorithms/math/FastPowering"
	"github.com/gongchengra/learning-go/02.algorithms/math/Fibonacci"
	"github.com/gongchengra/learning-go/02.algorithms/math/IsPowerOfTwo"
	"github.com/gongchengra/learning-go/02.algorithms/math/PascalTriangle"
	"github.com/gongchengra/learning-go/02.algorithms/math/PrimalityTest"
	"github.com/gongchengra/learning-go/02.algorithms/math/Radian"
	"github.com/gongchengra/learning-go/02.algorithms/math/SieveOfEratosthenes"
	"github.com/gongchengra/learning-go/02.algorithms/math/SquareRoot"
)

func main() {
	fmt.Println(EuclideanAlgorithm.GCD(10, 20))
	fmt.Println(Factorial.Factorial(18))
	fmt.Println(Factorial.FactorialRecursive(20))
	fmt.Println(FastPowering.FastPowering(3.0, 5))
	fmt.Println(Fibonacci.FibonacciRecursive(10))
	fmt.Println(Fibonacci.FibonacciSequence(10))
	fmt.Println(IsPowerOfTwo.IsPowerOfTwo(10))
	fmt.Println(IsPowerOfTwo.IsPowerOfTwoBitwise(1024))
	fmt.Println(PascalTriangle.PascalTriangle(10))
	fmt.Println(PrimalityTest.IsPrime(9))
	fmt.Println(Radian.DegreeToRadian(10))
	fmt.Println(Radian.RadianToDegree(math.Pi))
	fmt.Println(SieveOfEratosthenes.SieveOfEratosthenes(1000))
	fmt.Println(SquareRoot.SquareRoot(169))
}
