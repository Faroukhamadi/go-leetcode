package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	stack := make([]byte, 0)
	curr := ""

	var backtrack func()
	backtrack = func() {
		if len(curr) == n*2 {
			if len(stack) == 0 {
				res = append(res, curr)
			}
			return
		}
		if len(stack) > 0 && stack[len(stack)-1] == '(' {
			curr += ")"
			stack = stack[:len(stack)-1]
			backtrack()
			curr = curr[:len(curr)-1]
			stack = append(stack, '(')
		}
		curr += "("
		stack = append(stack, '(')
		backtrack()
		stack = stack[:len(stack)-1]
		curr = curr[:len(curr)-1]
	}
	backtrack()
	return res
}

func main() {

}
