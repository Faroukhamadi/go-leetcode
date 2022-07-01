package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			l := len(stack)

			if l == 0 {
				return false
			} else {
				pop := stack[l-1]
				stack = stack[:l-1]

				if pop != c {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValid("([[)"))
}
