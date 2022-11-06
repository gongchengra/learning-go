package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	bitset := make(map[byte]bool)
	res, l, r := 0, 0, 0
	for l < len(s) {
		if _, ok := bitset[s[r]]; ok {
			bitset[s[l]] = false
			l++
		} else {
			bitset[s[r]] = true
			r++
		}
		if res < r-l {
			res = r - l
		}
		if l+r >= len(s) || r >= len(s) {
			break
		}
	}
	return res
}
