package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("({}[]"))
}

// method 1

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := []rune{}
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

/* method 2

func isValid(str string) bool {
	s := new(stack)
	for _, b := range str {
		switch b {
		case '(', '[', '{':
			s.push(b)
		case ')', ']', '}':
			if r, ok := s.pop(); !ok || r != matching[b] {
				// !ok 说明“（[{”的数量，小于")]}"的数量
				return false
			}
		}
	}
	// len(*s) > 0 说明"([{"的数量，大于")]}"的数量
	if len(*s) > 0 {
		return false
	}
	return true
}
var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}

func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}
*/
