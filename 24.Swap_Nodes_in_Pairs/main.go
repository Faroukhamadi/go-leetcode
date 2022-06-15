package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	temp := &ListNode{0, head}
	cur := temp

	for cur.Next != nil && cur.Next.Next != nil {
		firstNode := cur.Next
		secondNode := cur.Next.Next
		firstNode.Next = secondNode.Next
		cur.Next = secondNode
		cur.Next.Next = firstNode
		cur = cur.Next.Next
	}

	return temp.Next
}

func main() {

}
