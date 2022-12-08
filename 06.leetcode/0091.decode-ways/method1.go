package main

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	result := 0
	if s[0] != '0' {
		result += numDecodings(s[1:])
	}
	//     if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
	if (s[0]-'0')*10+(s[1]-'0') <= 26 {
		result += numDecodings(s[2:])
	}
	return result
}
