package main

import "fmt"

type item struct {
	char byte
	freq int
}

func frequencySort(s string) string {
	freqMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		char := s[i]
		freqMap[char]++
	}

	var res string

	var heap []item

	for char, freq := range freqMap {
		heap = append(heap, item{
			char,
			freq,
		})
	}

	for i := len(heap)/2 - 1; i > -1; i-- {
		heapDown(heap, i, len(heap)-1)
	}

	for i := len(heap) - 1; i > -1; i-- {
		for j := 0; j < heap[0].freq; j++ {
			res += fmt.Sprintf("%c", heap[0].char)
		}

		heap[0], heap[i] = heap[i], heap[0]
		heapDown(heap, 0, i-1)
	}

	return res
}

func heapDown(arr []item, pos int, limit int) {
	l := 2*pos + 1
	r := 2*pos + 2

	larger := pos

	if l <= limit && arr[l].freq > arr[larger].freq {
		larger = l
	}

	if r <= limit && arr[r].freq > arr[larger].freq {
		larger = r
	}

	if larger != pos {
		arr[larger], arr[pos] = arr[pos], arr[larger]
		heapDown(arr, larger, limit)
	}

}

func main() {

}
