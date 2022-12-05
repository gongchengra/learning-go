package main

// Time Limit Exceeded for
// s1 = "eebaacbcbcadaaedceaaacadccd" and  s2 = "eadcaacabaddaceacbceaabeccd"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	ls, rec := len(s1), make([]int, 256)
	for i := 0; i < ls; i++ {
		rec[s1[i]]++
		rec[s2[i]]--
	}
	for i := range rec {
		if rec[i] != 0 {
			return false
		}
	}
	for i := 1; i < ls; i++ {
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[0:i], s2[ls-i:]) && isScramble(s1[i:], s2[0:ls-i]) {
			return true
		}
	}
	return false
}
