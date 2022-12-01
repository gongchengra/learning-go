package main

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	need := [256]int{}
	for i := range t {
		need[t[i]]++
	}
	has := [256]int{}
	min := ls + 1
	begin, end, wbegin, wend, count := 0, 0, 0, 0, 0
	for ; end < ls; end++ {
		if need[s[end]] == 0 {
			continue
		}
		if has[s[end]] < need[s[end]] {
			count++
		}
		has[s[end]]++
		if count == lt {
			for need[s[begin]] == 0 || has[s[begin]] > need[s[begin]] {
				if has[s[begin]] > need[s[begin]] {
					has[s[begin]]--
				}
				begin++
			}
			tmp := end - begin + 1
			if min > tmp {
				min, wbegin, wend = tmp, begin, end
			}
		}
	}
	if count < lt {
		return ""
	}
	return s[wbegin : wend+1]
}
