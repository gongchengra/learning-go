package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s /path/to/your/folder\n", os.Args[0])
		os.Exit(1)
	}
	folderPath := os.Args[1]
	fileHashes := make(map[string][]string)
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			hash := md5.New()
			if _, err := io.Copy(hash, file); err != nil {
				return err
			}
			hashInBytes := hash.Sum(nil)
			hashString := fmt.Sprintf("%x", hashInBytes)
			fileHashes[hashString] = append(fileHashes[hashString], path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through files:", err)
		os.Exit(2)
	}
	duplicatesFound := false
	fmt.Println("=== Duplicate file groups ===")
	for _, files := range fileHashes {
		if len(files) > 1 {
			duplicatesFound = true
			fmt.Printf("Group of files with hash %x:\n", md5.Sum([]byte(files[0])))
			for _, file := range files {
				fmt.Println(file)
			}
			fmt.Println()
		}
	}
	if !duplicatesFound {
		fmt.Println("No duplicate files found.")
		os.Exit(0)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Would you like to delete duplicates, keeping only one copy of each file? (yes/no) ")
	deleteChoice, _ := reader.ReadString('\n')
	deleteChoice = strings.TrimSpace(deleteChoice)
	if deleteChoice != "yes" {
		fmt.Println("Exiting without deleting any files.")
		os.Exit(0)
	}
	for hash, files := range fileHashes {
		if len(files) > 1 {
			fmt.Printf("Duplicates for hash %x:\n", md5.Sum([]byte(files[0])))
			fmt.Println("0) Keep all duplicates")
			for i, file := range files {
				fmt.Printf("%d) %s\n", i+1, file)
			}
			var keepChoice int
			for {
				fmt.Print("Enter the number of the file you wish to keep: ")
				_, err := fmt.Fscan(reader, &keepChoice)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error reading input:", err)
					continue
				}
				if keepChoice == 0 {
					fmt.Printf("Keeping all duplicate files for hash %x.\n", hash)
					break
				} else if keepChoice > 0 && keepChoice <= len(files) {
					keepIndex := keepChoice - 1
					for i, file := range files {
						if i != keepIndex {
							fmt.Println("Deleting:", file)
							err := os.Remove(file)
							if err != nil {
								fmt.Fprintf(os.Stderr, "Failed to delete file %s: %v\n", file, err)
							}
						}
					}
					fmt.Println("Kept:", files[keepIndex])
					break
				} else {
					fmt.Println("Invalid selection. Please choose a valid file number.")
				}
			}
			fmt.Println()
		}
	}
	fmt.Println("Duplicate files deletion complete.")
}
