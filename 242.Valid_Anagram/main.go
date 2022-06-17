package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		mp[t[i]]--
	}
	for key, val := range mp {
		fmt.Printf("key: %d, val: %d\n", key, val)
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v\n", isAnagram("anagram", "aagaram"))
}
