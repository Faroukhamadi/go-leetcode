package main

import "math"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, max(piles)
	res := r

	for l <= r {
		k := (l + r) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}
		if hours <= h {
			res = min(res, k)
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
}
