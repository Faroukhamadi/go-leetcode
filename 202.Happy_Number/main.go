package main

import "log"

func isHappy(n int) bool {
	switch n {
	case 1, 7:
		return true
	case 0, 2, 3, 4, 5, 6, 8, 9:
		return false
	default:
		return isHappy(sumNum(n))
	}
}

func sumNum(n int) int {
	out, last := 0, 0
	for n > 0 {
		last = n % 10
		out += last * last
		n /= 10
	}
	return out
}

func main() {
	log.Println(isHappy(19))
}
