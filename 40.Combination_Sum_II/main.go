package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := [][]int{}
	findCombinations(candidates, target, 0, nil, &ans)

	return ans
}

func findCombinations(candidates []int, target, idx int, temp []int, ans *[][]int) {
	if target == 0 {
		t := append([]int{}, temp...)
		*ans = append(*ans, t)
		return
	}

	if target < 0 {
		return
	}

	for i := idx; i < len(candidates); i++ {
		if i == idx || candidates[i] != candidates[i-1] {
			findCombinations(candidates, target-candidates[i], i+1, append(temp, candidates[i]), ans)
		}
		continue
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
