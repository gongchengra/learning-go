package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	b := []byte(s)
	m := make(map[byte][]int)
	for k, c := range b {
		m[c] = append(m[c], k)
	}
	max, l, r := 0, 0, 0
	for _, v := range m {
		for i := 0; i < len(v)-1; i++ {
			for j := len(v) - 1; j > i; j-- {
				if isPalindrome(b[v[i] : v[j]+1]) {
					if max < v[j]+1-v[i] {
						max, l, r = v[j]+1-v[i], v[i], v[j]
					}
				}
			}
		}
	}
	return string(b[l : r+1])
}

func isPalindrome(s []byte) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
