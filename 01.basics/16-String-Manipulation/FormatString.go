package main

// https://zetcode.com/golang/string-format/
/*
The functions take the format string and the list of arguments as parameters.

%[flags][width][.precision]verb

d - decimal integer
o - octal integer
O - octal integer with 0o prefix
b - binary integer
x - hexadecimal integer lowercase
X - hexadecimal integer uppercase
f - decimal floating point, lowercase
F - decimal floating point, uppercase
e - scientific notation (mantissa/exponent), lowercase
E - scientific notation (mantissa/exponent), uppercase
g - the shortest representation of %e or %f
G - the shortest representation of %E or %F
c - a character represented by the corresponding Unicode code point
q - a quoted character
U - Unicode escape sequence
t - the word true or false
s - a string
v - default format
#v - Go-syntax representation of the value
T - a Go-syntax representation of the type of the value
p - pointer address
% - a double %% prints a single %
*/

import "fmt"

type User struct {
	name       string
	occupation string
}

func main() {
	{
		name := "Jane"
		age := 17
		fmt.Printf("%s is %d years old\n", name, age)
		res := fmt.Sprintf("%s is %d years old", name, age)
		fmt.Println(res)
	}
	{
		msg := "and old falcon"
		n := 16
		w := 12.45
		r := true
		u := User{"John Doe", "gardener"}
		vals := []int{1, 2, 3, 4, 5}
		ctrs := map[string]string{
			"sk": "Slovakia",
			"ru": "Russia",
			"de": "Germany",
			"no": "Norway",
		}
		fmt.Printf("%v %v %v %v %v\n  %v %v\n", msg, n, w, u, r, vals, ctrs)
		fmt.Printf("%v %+v\n", u, u)
		fmt.Println("--------------------")
		fmt.Printf("%#v %#v %#v %#v %#v\n  %#v %#v\n", msg, n, w, u, r, vals, ctrs)
		fmt.Printf("%T %T %T %T %T %T %T\n", msg, n, w, u, r, vals, ctrs)
		fmt.Println("--------------------")
		fmt.Printf("The prices dropped by 12%%\n")
	}
	{
		n1 := 2
		n2 := 3
		n3 := 4
		res := fmt.Sprintf("There are %d oranges %d apples %d plums", n1, n2, n3)
		fmt.Println(res)
		res2 := fmt.Sprintf("There are %[2]d oranges %d apples %[1]d plums", n1, n2, n3)
		fmt.Println(res2)
	}
	{
		fmt.Printf("%d\n", 1671)
		fmt.Printf("%o\n", 1671)
		fmt.Printf("%x\n", 1671)
		fmt.Printf("%X\n", 1671)
		fmt.Printf("%#b\n", 1671)
		fmt.Printf("%f\n", 1671.678)
		fmt.Printf("%F\n", 1671.678)
		fmt.Printf("%e\n", 1671.678)
		fmt.Printf("%E\n", 1671.678)
		fmt.Printf("%g\n", 1671.678)
		fmt.Printf("%G\n", 1671.678)
		fmt.Printf("%s\n", "Zetcode")
		fmt.Printf("%c %c %c %c %c %c %c\n", 'Z', 'e', 't', 'C', 'o', 'd', 'e')
		fmt.Printf("%p\n", []int{1, 2, 3})
		fmt.Printf("%d%%\n", 1671)
		fmt.Printf("%t\n", 3 > 5)
		fmt.Printf("%t\n", 5 > 3)
		fmt.Printf("%0.f\n", 16.540)
		fmt.Printf("%0.2f\n", 16.540)
		fmt.Printf("%0.3f\n", 16.540)
		fmt.Printf("%0.5f\n", 16.540)
	}
	{
		val := 122
		fmt.Printf("%d\n", val)
		fmt.Printf("%c\n", val)
		fmt.Printf("%q\n", val)
		fmt.Printf("%x\n", val)
		fmt.Printf("%X\n", val)
		fmt.Printf("%o\n", val)
		fmt.Printf("%O\n", val)
		fmt.Printf("%b\n", val)
		fmt.Printf("%U\n", val)
	}
	{
		val := 1273.78888769000
		fmt.Printf("%f\n", val)
		fmt.Printf("%e\n", val)
		fmt.Printf("%g\n", val)
		fmt.Printf("%E\n", val)
		fmt.Printf("%G\n", val)
		fmt.Println("-------------------------")
		fmt.Printf("%.10f\n", val)
		fmt.Printf("%.10e\n", val)
		fmt.Printf("%.10g\n", val)
		fmt.Printf("%.10E\n", val)
		fmt.Printf("%.10G\n", val)
		fmt.Println("-------------------------")
		val2 := 66_000_000_000.1200
		fmt.Printf("%f\n", val2)
		fmt.Printf("%e\n", val2)
		fmt.Printf("%g\n", val2)
		fmt.Printf("%E\n", val2)
		fmt.Printf("%G\n", val2)
	}
	{
		fmt.Printf("% d\n", 1691)
		fmt.Printf("%+d\n", 1691)
		fmt.Println("---------------------")
		fmt.Printf("%#x\n", 1691)
		fmt.Printf("%#X\n", 1691)
		fmt.Printf("%#b\n", 1691)
		fmt.Println("---------------------")
		fmt.Printf("%10d\n", 1691)
		fmt.Printf("%-10d\n", 1691)
		fmt.Printf("%010d\n", 1691)
	}
	{
		w := "falcon"
		n := 122
		h := 455.67
		fmt.Printf("%s\n", w)
		fmt.Printf("%10s\n", w)
		fmt.Println("---------------------")
		fmt.Printf("%d\n", n)
		fmt.Printf("%7d\n", n)
		fmt.Printf("%07d\n", n)
		fmt.Println("---------------------")
		fmt.Printf("%10f\n", h)
		fmt.Printf("%11f\n", h)
		fmt.Printf("%12f\n", h)
	}
}
