package main

import (
	"fmt"
)

// O(n) time complexity
func containsDuplicate(nums []int) bool {
	hashSet := make(map[int]bool)
	for index, item := range nums {
		fmt.Printf("index: %d\n", index)
		if hashSet[item] {
			return true
		} else {
			hashSet[item] = true
		}
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 6}
	fmt.Println(containsDuplicate(nums))
}
