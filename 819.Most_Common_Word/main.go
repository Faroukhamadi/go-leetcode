package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

// 1.18 sol not compatible with leetcode :(
func mostCommonWord(paragraph string, banned []string) string {
	bannedWords := map[string]bool{}
	for _, word := range banned {
		bannedWords[word] = true
	}

	counts := map[string]int{}
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	for _, word := range strings.Split(strings.ToLower(reg.ReplaceAllString(paragraph, " ")), " ") {
		if !slices.Contains(banned, word) {
			counts[word]++
		}
	}

	result := ""
	for k, v := range counts {
		if result == "" || v > counts[result] {
			result = k
		}
	}

	return result
}

func main() {
	fmt.Println(mostCommonWord("a.", []string{}))
}
