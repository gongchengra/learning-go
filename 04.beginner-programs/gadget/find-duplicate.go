package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
			hash, err := hashFileMd5(path)
			if err != nil {
				return err
			}
			fileHashes[hash] = append(fileHashes[hash], path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through files:", err)
		os.Exit(1)
	}

	duplicatesFound := false
	fmt.Println("=== Duplicate file groups ===")
	for hash, files := range fileHashes {
		if len(files) > 1 {
			duplicatesFound = true
			fmt.Printf("Group of files with hash %s:\n", hash)
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

	var deleteChoice string
	fmt.Print("Would you like to delete duplicates, keeping only one copy of each file? (yes/no) ")
	fmt.Scan(&deleteChoice)

	if deleteChoice != "yes" {
		fmt.Println("Exiting without deleting any files.")
		os.Exit(0)
	}

	for hash, files := range fileHashes {
		if len(files) > 1 {
			fmt.Printf("Duplicates for hash %s:\n", hash)
			for i, file := range files {
				fmt.Printf("%d) %s\n", i+1, file)
			}

			var keepChoice int
			for {
				fmt.Print("Enter the number of the file you wish to keep: ")
				_, err := fmt.Scan(&keepChoice)
				if err != nil {
					fmt.Println("Invalid input. Please enter a number.")
					continue
				}
				if keepChoice > 0 && keepChoice <= len(files) {
					for i, file := range files {
						if i != keepChoice-1 {
							fmt.Println("Deleting:", file)
							err := os.Remove(file)
							if err != nil {
								fmt.Println("Error deleting file:", err)
							}
						}
					}
					fmt.Println("Kept:", files[keepChoice-1])
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

func hashFileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
