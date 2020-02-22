package main

import "fmt"

func isValid(s string) bool {
	chs := make([]int32, 1000)
	index := 0 // index 总是指向空节点
	for _, ch := range s {
		switch ch {
		case '[':
			fallthrough
		case '(':
			fallthrough
		case '{':
			chs[index] = ch
			index++
		case ']':
			fallthrough
		case ')':
			fallthrough
		case '}':
			if index == 0 {
				return false
			}
			index--
			c := chs[index]
			if ch == ']' && c != '[' {
				return false
			}
			if ch == ')' && c != '(' {
				return false
			}
			if ch == '}' && c != '{' {
				return false
			}
		default:
			continue
		}
	}
	if index > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("([)]"))
}
