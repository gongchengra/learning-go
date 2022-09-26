package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
)

type User struct {
	Id         string `csv:"user_id"`
	Name       string `csv:"user_name"`
	Occupation string `csv:"user_occupation"`
}

func (p User) String() string {
	return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
		p.Id, p.Name, p.Occupation)
}

func main() {
	users := []User{}
	users = append(users, User{Id: "1", Name: "John Doe", Occupation: "gardener"})
	users = append(users, User{Id: "2", Name: "Roger Doe", Occupation: "driver"})
	res, _ := gocsv.MarshalString(users)
	fmt.Println(res)
}
