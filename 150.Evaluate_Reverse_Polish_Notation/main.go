package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	num, _ := strconv.Atoi(tokens[0])
	stack = append(stack, num)
	i := 1
	for i < len(tokens) {
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, num)
		} else {
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, l+r)
			case "-":
				stack = append(stack, l-r)
			case "*":
				stack = append(stack, l*r)
			case "/":
				stack = append(stack, l/r)
			}
		}
		i++
	}

	return stack[0]
}

func main() {

}
