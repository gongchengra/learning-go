package main

import (
	"bufio"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return err
}

func main() {
	fs.WalkDir(os.DirFS("."), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		ext := filepath.Ext(d.Name())
		if d.IsDir() || ext != ".go" {
			return nil
		}
		in, err := os.Open(path)
		if err != nil {
			fmt.Println("open file fail:", err)
			os.Exit(-1)
		}
		defer in.Close()
		out := ""
		br := bufio.NewReader(in)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("read err:", err)
				os.Exit(-1)
			}
			emptyLINE := regexp.MustCompile(`^\s*$`)
			if !emptyLINE.MatchString(string(line)) {
				f := regexp.MustCompile(`^func`)
				fs := f.ReplaceAllString(string(line), "\nfunc")
				t := regexp.MustCompile(`^type`)
				rs := t.ReplaceAllString(fs, "\ntype")
				out += rs + "\n"
			}
		}
		content, err := format.Source([]byte(out))
		if err != nil {
			return err
		}
		err = WriteToFile(path, string(content))
		if err != nil {
			return err
		}
		return nil
	})
}
