package main

import "fmt"

func reverse(x int) int {
	rev := 0

	for ; x != 0; x /= 10 {
		rev = rev*10 + x%10
	}
	if rev > -2147483647 && rev < 2147483646 {
		return rev
	}
	return 0
}

func main() {
	fmt.Println(reverse(7502))
}
