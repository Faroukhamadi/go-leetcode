package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		val, present := mp[diff]
		if present {
			return []int{val, i}
		}
		mp[n] = i
	}
	return nil
}

func main() {

}
