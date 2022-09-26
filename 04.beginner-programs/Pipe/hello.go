package main

// https://zetcode.com/golang/pipe/
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Hello %s!\n", buf)
	} else {
		fmt.Print("Enter your name: ")
		var name string
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello %s!\n", name)
	}
}
