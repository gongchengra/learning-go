package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	check := func(s string) bool {
		if i, _ := strconv.Atoi(s); i > 255 {
			return false
		}
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		return true
	}
	var res, t []string
	var dfs func(s string, t []string)
	dfs = func(s string, t []string) {
		if len(t) == 4 {
			if s == "" {
				res = append(res, strings.Join(t, "."))
			}
			return
		}
		for i := 1; i < 4 && i <= len(s); i++ {
			if check(s[0:i]) {
				t = append(t, s[0:i])
				dfs(s[i:], t)
				t = t[0 : len(t)-1]
			}
		}
	}
	dfs(s, t)
	return res
}
