package main

import (
	"fmt"
	"laogongshuo.com/string/HammingDistance"
	"laogongshuo.com/string/LevenshteinDistance"
)

func main() {
	fmt.Println(HammingDistance.HammingDistance("kitten", "sitten"))
	fmt.Println(LevenshteinDistance.LevenshteinDistance("abc", "acdef"))
}
