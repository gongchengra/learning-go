package main

/*
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts
byte        alias for uint8
rune        alias for int32
uint     unsigned, either 32 or 64 bits
int      signed, either 32 or 64 bits
uintptr  unsigned integer large enough to store the uninterpreted bits of a pointer value
*/
import (
	"fmt"
	"math"
)

func main() {
	// Integer
	var i int = 5
	fmt.Println(i)
	// String
	var s string = "Hello World!"
	fmt.Println(s)
	// Float
	var f float64 = 3.14159265359
	fmt.Println(f)
	// Boolean
	var b bool = true
	fmt.Println(b)
	b = false
	fmt.Println(b)
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	// Usually a Printf format string containing multiple % verbs
	// would require the same number of extra operands, but the [1] ‘‘adverbs’ af ter % te ll Printf to
	// use the first operand over and over again. Second, the # adverb for %o or %x or %X tells Printf to emit a 0 or 0x or 0X prefix respectively.
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	fmt.Printf("%d %d %d %[1]c %[2]c %[1]q %[2]q %[3]q\n", ascii, unicode, newline)
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
