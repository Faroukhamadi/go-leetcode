package main

import "fmt"

func subsets(nums []int) [][]int {
	var subsets [][]int
	generateSubsets(0, nums, []int{}, &subsets)

	return subsets
}

func generateSubsets(index int, nums []int, current []int, subsets *[][]int) {
	currentCopy := make([]int, len(current))
	copy(currentCopy, current)
	*subsets = append(*subsets, currentCopy)
	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])
		generateSubsets(i+1, nums, current, subsets)
		current = current[:len(current)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
