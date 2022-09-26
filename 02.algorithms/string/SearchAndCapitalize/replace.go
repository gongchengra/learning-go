package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("lack of arguments, eg: replace ${stringToCapitalize} ${path_of_file}")
		os.Exit(-1)
	}
	fileName := os.Args[2]
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, os.Args[1]) {
			lines[i] = strings.Replace(lines[i], os.Args[1], strings.Title(os.Args[1]), -1)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
