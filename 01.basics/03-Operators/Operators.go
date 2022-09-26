package main

import "fmt"

func main() {
	var i int = 10
	var k int = 20
	var z int = 30
	// Arithmetic Operators
	fmt.Printf("i + k = %d\n", i+k)
	fmt.Printf("i - k = %d\n", i-k)
	fmt.Printf("i * k = %d\n", i*k)
	fmt.Printf("i / k = %d\n", i/k)
	fmt.Printf("i mod k = %d\n", i%k)
	// Comparison Operators
	fmt.Println(i == k)
	fmt.Println(i != k)
	fmt.Println(i < k)
	fmt.Println(i <= k)
	fmt.Println(i > k)
	fmt.Println(i >= k)
	// Logical Operators
	fmt.Println(i < z && i > k)
	fmt.Println(i < z || i > k)
	fmt.Println(!(i == z && i > k))
	// Assignment Operators
	var x, y = 15, 25
	x = y
	fmt.Println("= ", x)
	x = 15
	x += y
	fmt.Println("+=", x)
	x = 50
	x -= y
	fmt.Println("-=", x)
	x = 2
	x *= y
	fmt.Println("*=", x)
	x = 100
	x /= y
	fmt.Println("/=", x)
	x = 40
	x %= y
	fmt.Println("%=", x)
	{
		var x uint8 = 1<<1 | 1<<5
		var y uint8 = 1<<1 | 1<<2
		fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
		fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
		fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
		fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
		fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
		fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
		// z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1;
		// other wise it equals the corresponding bit of x.
		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 { // membership test
				fmt.Println(i) // "1", "5"
			}
		}
		fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
		fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
	}
}
