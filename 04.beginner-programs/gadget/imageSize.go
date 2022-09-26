package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		ext := filepath.Ext(info.Name())
		if info.IsDir() || (ext != "gif" && ext != "png" && ext != "jpg") {
			return nil
		}
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		image, _, err := image.DecodeConfig(file) // Image Struct
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", path, err)
		}
		fmt.Println("File:", path, "Width:", image.Width, "Height:", image.Height)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
