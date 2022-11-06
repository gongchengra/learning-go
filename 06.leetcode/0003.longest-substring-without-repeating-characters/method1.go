package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	bytes := []byte(s)
	max := 1
	for i := 0; i < len(bytes); i++ {
		c := make(map[byte]bool)
		c[bytes[i]] = true
		for j := i + 1; j < len(bytes); j++ {
			if _, ok := c[bytes[j]]; ok {
				if max < len(c) {
					max = len(c)
				}
				break
			} else {
				c[bytes[j]] = true
				if max < len(c) {
					max = len(c)
				}
			}
		}
	}
	return max
}
