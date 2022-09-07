package main

import "fmt"

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}

	mapping := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	letterCombinationsRecursive(&res, digits, "", 0, mapping)

	return res
}

func letterCombinationsRecursive(res *[]string, digits, current string, index int, mapping []string) {
	if index == len(digits) {
		*res = append(*res, current)
		return
	}

	letters := mapping[digits[index]-'0']
	for i := 0; i < len(letters); i++ {
		letterCombinationsRecursive(res, digits, current+string(letters[i]), index+1, mapping)
	}

}

func main() {
	fmt.Println(letterCombinations("23"))
}
