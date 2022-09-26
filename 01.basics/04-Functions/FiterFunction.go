package main

import "fmt"

type User struct {
	name       string
	occupation string
	married    bool
}

func main() {
	u1 := User{"John Doe", "gardener", false}
	u2 := User{"Richard Roe", "driver", true}
	u3 := User{"Bob Martin", "teacher", true}
	u4 := User{"Lucy Smith", "accountant", false}
	u5 := User{"James Brown", "teacher", true}
	users := []User{u1, u2, u3, u4, u5}
	married := filter(users, func(u User) bool { return u.married == true })
	fmt.Println("Married:")
	fmt.Printf("%v\n", married)
	teachers := filter(users, func(u User) bool { return u.occupation == "teacher" })
	fmt.Println("Teachers:")
	fmt.Printf("%v\n", teachers)
	mt := filter(users, func(u User) bool { return u.occupation == "teacher" && u.married == true })
	fmt.Println("Married teachers:")
	fmt.Printf("%v\n", mt)
	driver := func(s []User) []User {
		var res []User
		for _, v := range s {
			if v.occupation == "driver" {
				res = append(res, v)
			}
		}
		return res
	}(users)
	fmt.Println("Drivers:")
	fmt.Println(driver)
	gardener := func(s []User) []User {
		var res []User
		for _, v := range s {
			if v.occupation == "gardener" {
				res = append(res, v)
			}
		}
		return res
	}
	fmt.Println("Gardener:")
	fmt.Println(gardener(users))
	of := func(s []User, f string) []User {
		var res []User
		for _, v := range s {
			if v.occupation == f {
				res = append(res, v)
			}
		}
		return res
	}
	fmt.Println("Occupation Filter Function: of(users, \"accountant\"):")
	fmt.Println(of(users, "accountant"))
}

func filter(s []User, f func(User) bool) []User {
	var res []User
	for _, v := range s {
		if f(v) == true {
			res = append(res, v)
		}
	}
	return res
}
