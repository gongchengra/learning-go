package main

import (
	"fmt"
	"github.com/gongchengra/learning-go/02.algorithms/math/Fibonacci"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(Fibonacci.FibonacciRecursive(42))      //     267914296
	fmt.Println("seconds it took:", time.Since(start)) //     seconds it took: 1.443081365s
	start = time.Now()
	res := Fibonacci.FibonacciSequence(42)
	fmt.Printf("%9.f\n", res[42])                      //     267914296
	fmt.Println("seconds it took:", time.Since(start)) //     seconds it took: 22.701µs
	res = Fibonacci.FibonacciSequence(1476)            // the max value float64 could hold
	fmt.Printf("%9.f\n", res[len(res)-1])
	//     130698922376339873754511593703999304853661815941920982715896371280424691495866567130509827216117625177952738381240755518030797439683443697785696230802473309617042775347304891963181519627287463521203531259388682404883801028462229399345567884825464934136563115441584430300333788777345438315116223032518554681344
	fmt.Println("seconds it took:", time.Since(start)) //     seconds it took: 141.458µs
}
