package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	index := 0
	i := 0
	for i < len(chars) {
		j := i
		for j < len(chars) && chars[j] == chars[i] {
			j++
		}

		chars[index] = chars[i]
		index++
		if j-i > 1 {
			count := strconv.Itoa(j - i)
			for _, c := range count {
				chars[index] = byte(c)
				index++
			}
		}
		i = j
	}
	return index
}

func main() {
	fmt.Println(compress([]byte("aabbccc")))
}
