package Fibonacci

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	} else {
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}
}
