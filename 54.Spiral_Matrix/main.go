package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var nums []int

	top := 0
	bot := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	size := len(matrix) * len(matrix[0])

	for len(nums) < size {
		for i := left; i <= right && len(nums) < size; i++ {
			nums = append(nums, matrix[top][i])
		}
		top++
		for i := top; i <= bot && len(nums) < size; i++ {
			nums = append(nums, matrix[i][right])
		}
		right--
		for i := right; i >= left && len(nums) < size; i-- {
			nums = append(nums, matrix[bot][i])
		}
		bot--
		for i := bot; i >= top && len(nums) < size; i-- {
			nums = append(nums, matrix[i][left])
		}
		left++
	}
	return nums
}

func main() {
	nums := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println("nums: ", nums)
}
