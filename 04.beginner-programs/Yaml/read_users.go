package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type User struct {
	Name       string
	Occupation string
}

func main() {
	yfile, err := ioutil.ReadFile("users.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]User)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}
