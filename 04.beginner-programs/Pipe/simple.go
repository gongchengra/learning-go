package main

// https://zetcode.com/golang/pipe/
import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	{
		r, w := io.Pipe()
		go func() {
			fmt.Fprint(w, "Hello there\n")
			w.Close()
		}()
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
