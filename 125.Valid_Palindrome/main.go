package main

import "strings"

func isPalindrome(s string) bool {
	n := len(s)
	s = strings.ToLower(s)

	l := 0
	r := n - 1

	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return ((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'))
}
