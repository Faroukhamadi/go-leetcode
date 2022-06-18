package main

func topKFrequent(nums []int, k int) []int {
	dic := map[int]int{}
	for _, n := range nums {
		dic[n]++
	}
	revdic := map[int][]int{}
	for n, count := range dic {
		revdic[count] = append(revdic[count], n)
	}
	out := []int{}
	for i := len(nums); len(out) != k; i-- {
		for _, n := range revdic[i] {
			if len(out) != k {
				out = append(out, n)
			}
		}
	}
	return out
}

func main() {

}
