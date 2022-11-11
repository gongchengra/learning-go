package main

import "fmt"

func main() {
	//     fmt.Println(longestValidParentheses("(()"))
	//     fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()(())))"))
}

func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
		fmt.Println(s[i], stack, res)
	}
	return res
}

/*

func longestValidParentheses(s string) int {
	l, r, maxlen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			maxlen = max(maxlen, 2*r)
		} else if r > l {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			maxlen = max(maxlen, 2*l)
		} else if l > r {
			l, r = 0, 0
		}
	}
	return maxlen
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
