package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s /path/to/your/folder\n", os.Args[0])
		os.Exit(1)
	}

	folderPath := os.Args[1]
	hashes := make(map[string][]string)

	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			hash, err := fileMD5(path)
			if err != nil {
				return err
			}
			hashes[hash] = append(hashes[hash], path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through directory:", err)
		os.Exit(2)
	}

	duplicatesFound := false
	fmt.Println("=== Duplicate file groups ===")
	for hash, files := range hashes {
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

	for hash, files := range hashes {
		if len(files) > 1 {
			fmt.Printf("Duplicates for hash %s:\n", hash)
			fmt.Println("0) Keep all duplicates")
			for i, file := range files {
				fmt.Printf("%d) %s\n", i+1, file)
			}

			var keepChoice string
			for {
				fmt.Print("Enter the number of the file you wish to keep, or 'd' to delete all: ")
				fmt.Scan(&keepChoice)

				if keepChoice == "d" {
					for _, file := range files {
						fmt.Printf("Deleting: %s\n", file)
						os.Remove(file)
					}
					fmt.Println("All duplicates have been deleted.")
					break
				} else if keep, err := strconv.Atoi(keepChoice); err == nil && keep == 0 {
					fmt.Printf("Keeping all duplicate files for hash %s.\n", hash)
					break
				} else if keep > 0 && keep <= len(files) {
					keepIndex := keep - 1
					for i, file := range files {
						if i != keepIndex {
							fmt.Printf("Deleting: %s\n", file)
							os.Remove(file)
						}
					}
					fmt.Printf("Kept: %s\n", files[keepIndex])
					break
				} else {
					fmt.Println("Invalid selection. Please choose a valid file number or 'd' to delete all.")
				}
			}
			fmt.Println()
		}
	}

	fmt.Println("Duplicate files deletion complete.")
}

func fileMD5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}
