package main

// https://zetcode.com/golang/pipe/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "webcode.me")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)
	num := 0
	for {
		line, _, _ := buf.ReadLine()
		if num > 3 {
			os.Exit(0)
		}
		num += 1
		fmt.Println(string(line))
	}
}
