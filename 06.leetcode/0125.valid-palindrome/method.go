package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; {
		if isAlphanumeric(s[i]) && isAlphanumeric(s[j]) {
			if s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
		} else if isAlphanumeric(s[i]) && !isAlphanumeric(s[j]) {
			j--
		} else {
			i++
		}
	}
	return true
}

func isAlphanumeric(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	} else {
		return false
	}
}
