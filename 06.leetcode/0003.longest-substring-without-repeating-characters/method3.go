package main

func lengthOfLongestSubstring(s string) int {
	location := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		location[s[i]] = -1
	}
	max, left := 0, 0
	for i := 0; i < len(s); i++ {
		if idx, _ := location[s[i]]; idx >= left {
			left = idx + 1
		} else if i+1-left > max {
			max = i + 1 - left
		}
		location[s[i]] = i
	}
	return max
}
