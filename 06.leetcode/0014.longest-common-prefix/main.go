package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	short := strs[0]
	shortlen := len(short)
	for k, v := range strs {
		if len(v) < shortlen {
			short = strs[k]
			shortlen = len(short)
		}
	}
	if shortlen == 0 {
		return ""
	}
	for i, c := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(c) {
				return strs[j][:i]
			}
		}
	}
	return short
}
