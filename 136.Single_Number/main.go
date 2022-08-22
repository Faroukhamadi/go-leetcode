package main

func singleNumber(nums []int) int {
	n := 0

	for _, v := range nums {
		n = n ^ v
	}

	return n
}
