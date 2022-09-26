package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	words := [5]string{"falcon", "sky", "earth", "cloud", "fox"}
	data, err := yaml.Marshal(&words)
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile("words.yaml", data, 0)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("data written")
}
