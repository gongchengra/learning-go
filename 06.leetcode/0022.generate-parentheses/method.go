package main

func generateParenthesis(n int) []string {
	res := []string{}
	ptr := &res
	help(ptr, n, n, "")
	return res
}

func help(res *[]string, left int, right int, str string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	}
	if left > 0 {
		help(res, left-1, right, str+"(")
	}
	if right > 0 && left < right {
		help(res, left, right-1, str+")")
	}
}
