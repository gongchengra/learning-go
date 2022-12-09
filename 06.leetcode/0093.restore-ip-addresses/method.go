package main

import (
	"strings"
)

func isValidIP(s string) bool {
	// 判断字符串 s 是否为合法的 IP 地址
	if len(s) == 0 || len(s) > 3 || (s[0] == '0' && len(s) > 1) {
		return false
	}
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num >= 0 && num <= 255
}

func restoreIpAddresses(s string) []string {
	var res []string
	n := len(s)
	if n < 4 || n > 12 {
		return res
	}
	for i := 1; i <= 3; i++ {
		for j := i + 1; j <= i+3; j++ {
			for k := j + 1; k <= j+3; k++ {
				if i >= n || j >= n || k >= n {
					continue
				}
				a, b, c, d := s[:i], s[i:j], s[j:k], s[k:]
				if isValidIP(a) && isValidIP(b) && isValidIP(c) && isValidIP(d) {
					res = append(res, strings.Join([]string{a, b, c, d}, "."))
				}
			}
		}
	}
	return res
}
