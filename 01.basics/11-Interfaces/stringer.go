package main

// https://zetcode.com/golang/interface/
/*
The Stringer interface is defined in the fmt package. Its String function is invoked when a type is passed to any of the print functions. We can customize the output message of our own types.

type Stringer interface {
    String() string
}
*/
import "fmt"

type User struct {
	Name       string
	Occupation string
}

func (u User) String() string {
	return fmt.Sprintf("%s is a(n) %s", u.Name, u.Occupation)
}

func main() {
	u1 := User{"John Doe", "gardener"}
	u2 := User{"Roger Roe", "driver"}
	fmt.Println(u1)
	fmt.Println(u2)
}
