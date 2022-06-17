package main

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	mp := map[[26]int][]string{}

	for _, str := range strs {
		encodedVal := encode(str)
		if mp[encodedVal] != nil {
			mp[encodedVal] = append(mp[encodedVal], str)
		} else {
			mp[encodedVal] = []string{str}
		}
	}

	for _, encodedStrArr := range mp {
		res = append(res, encodedStrArr)
	}
	return res
}

func encode(str string) [26]int {
	// use output array as the key for hashmap
	countArr := [26]int{}
	for _, char := range str {
		countArr[int(char)-'a']++
	}
	return countArr
}

func main() {

}
