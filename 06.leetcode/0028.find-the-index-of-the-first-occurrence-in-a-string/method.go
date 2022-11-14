package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for ih := 0; ih <= len(haystack)-len(needle); ih++ {
		found := 0
		if haystack[ih] == needle[0] {
			for k := 0; k < len(needle); k++ {
				if haystack[ih+k] == needle[k] {
					found++
				} else {
					break
				}
			}
			if found == len(needle) {
				return ih
			}
		}
	}
	return -1
}
