package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	Id         int    `xml:"id"`
	Name       string `xml:"name"`
	Occupation string `xml:"occupation"`
}

func (p User) String() string {
	return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
		p.Id, p.Name, p.Occupation)
}

func main() {
	user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
	res, _ := xml.MarshalIndent(user, " ", "  ")
	fmt.Println(xml.Header + string(res))
}
