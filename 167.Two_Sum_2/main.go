package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for true {
		sum := numbers[i] + numbers[j]
		switch {
		case sum == target:
			return []int{i + 1, j + 1}
		case sum > target:
			j--
		default:
			i++
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 3, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
