package main

import "fmt"

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1

	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func maxAreaBruteForce(height []int) int {
	// brute force
	res := 0

	for l := 0; l < len(height); l++ {
		for r := 0; r < len(height); r++ {
			area := (r - l) * min(height[l], height[r])
			res = max(res, area)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{}
	fmt.Println(maxArea(height))
}
