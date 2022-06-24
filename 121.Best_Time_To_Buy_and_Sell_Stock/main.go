package main

func maxProfit(prices []int) int {
	// left=buy, right=sell
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		// profitable ?
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}
	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nothing more annoying than gopls
	maxProfit([]int{})
}
