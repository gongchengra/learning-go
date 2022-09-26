package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))
	fmt.Println("On Unix:")
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
	fmt.Println("On Unix:")
	fmt.Println(filepath.IsAbs("/home/gopher"))
	fmt.Println(filepath.IsAbs(".bashrc"))
	fmt.Println(filepath.IsAbs(".."))
	fmt.Println(filepath.IsAbs("."))
	fmt.Println(filepath.IsAbs("/"))
	fmt.Println(filepath.IsAbs(""))
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
	fmt.Println(filepath.Join("a/b", "../../../xyz"))
	fmt.Println("On Unix:")
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))
	{
		paths := []string{
			"/a/b/c",
			"/b/c",
			"./b/c",
		}
		base := "/a"
		fmt.Println("On Unix:")
		for _, p := range paths {
			rel, err := filepath.Rel(base, p)
			fmt.Printf("%q: %q %v\n", p, rel, err)
		}
	}
	{
		paths := []string{
			"/home/arnie/amelia.jpg",
			"/mnt/photos/",
			"rabbit.jpg",
			"/usr/local//go",
		}
		fmt.Println("On Unix:")
		for _, p := range paths {
			dir, file := filepath.Split(p)
			fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
		}
	}
	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
	{
		tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
		if err != nil {
			fmt.Printf("unable to create test dir tree: %v\n", err)
			return
		}
		defer os.RemoveAll(tmpDir)
		os.Chdir(tmpDir)
		subDirToSkip := "skip"
		fmt.Println("On Unix:")
		err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			if info.IsDir() && info.Name() == subDirToSkip {
				fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
				return filepath.SkipDir
			}
			fmt.Printf("visited file or dir: %q\n", path)
			return nil
		})
		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
			return
		}
	}
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}
	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}
	return tmpDir, nil
}
