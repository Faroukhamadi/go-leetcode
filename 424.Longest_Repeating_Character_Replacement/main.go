package main

import "fmt"

func characterReplacement(s string, k int) int {
	windowStart, maxLength, maxRepeatedLetterCount := 0, 0, 0
	frequencyMap := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		rightChar := s[windowEnd]
		frequencyMap[rightChar]++

		maxRepeatedLetterCount = max(maxRepeatedLetterCount, frequencyMap[rightChar])

		if windowEnd-windowStart+1-maxRepeatedLetterCount > k {
			leftChar := s[windowStart]
			frequencyMap[leftChar]++
			windowStart++
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}
