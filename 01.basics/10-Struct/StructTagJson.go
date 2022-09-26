package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation,omitempty"`
}

func (p User) String() string {
	return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
		p.Id, p.Name, p.Occupation)
}

func main() {
	user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
	res, _ := json.MarshalIndent(user, " ", "  ")
	fmt.Println(string(res))
	user2 := User{Id: 1, Name: "John Doe"}
	res2, _ := json.MarshalIndent(user2, " ", "  ")
	fmt.Println(string(res2))
}
