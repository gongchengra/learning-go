package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Record struct {
	Item string `yaml:"item"`
	Col  string `yaml:"colour"`
	Size string `yaml:"size"`
}

type Config struct {
	Record Record `yaml:"Settings"`
}

func main() {
	config := Config{Record: Record{Item: "window", Col: "blue", Size: "small"}}
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile("config.yaml", data, 0)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("data written")
}
