package main

type KthLargest struct {
	k    int
	heap []int
}

func (this *KthLargest) add(val int) int {
	this.heap = append(this.heap, val)
	this.heapUp(len(this.heap) - 1)

	for len(this.heap) > this.k {
		this.pop()
	}

	return this.heap[0]
}

func (this *KthLargest) pop() int {
	if len(this.heap) == 0 {
		panic("empty heap")
	}

	poppedItem := this.heap[0]

	this.heap[0], this.heap[len(this.heap)-1] = this.heap[len(this.heap)-1], this.heap[0]
	this.heap = this.heap[:len(this.heap)-1]
	this.heapDown(0, len(this.heap)-1)

	return poppedItem
}

func (this *KthLargest) heapUp(pos int) {
	parent := (pos - 1) / 2

	if parent >= 0 && this.heap[pos] < this.heap[parent] {
		this.heap[pos], this.heap[parent] = this.heap[parent], this.heap[pos]
		this.heapUp(parent)
	}
}

func (this *KthLargest) heapDown(pos int, limit int) {
	l, r := 2*pos+1, 2*pos+2
	smaller := pos

	if l <= limit && this.heap[l] < this.heap[smaller] {
		smaller = l
	}

	if r <= limit && this.heap[r] < this.heap[smaller] {
		smaller = r
	}

	if smaller != pos {
		this.heap[smaller], this.heap[pos] = this.heap[pos], this.heap[smaller]
		this.heapDown(smaller, limit)
	}
}

func Constructor(k int, nums []int) KthLargest {
	res := KthLargest{
		k:    k,
		heap: nums,
	}

	for i := len(res.heap)/2 - 1; i > -1; i-- {
		res.heapDown(i, len(res.heap)-1)
	}

	return res
}

func (this *KthLargest) Add(val int) int {
	return this.add(val)
}
