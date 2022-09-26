package main

// modified from  https://pkg.go.dev/io/fs#example-WalkDir
import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fs.WalkDir(os.DirFS("."), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}
