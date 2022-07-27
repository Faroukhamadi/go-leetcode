package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMid(head *ListNode) *ListNode {
	var midPrev *ListNode = nil
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		midPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	var mid = midPrev.Next
	midPrev.Next = nil
	return mid
}

func merge(l1, l2 *ListNode) *ListNode {
	var head = &ListNode{0, nil}
	var cur = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

func main() {
	head := &ListNode{3, nil}
	second := &ListNode{2, nil}
	third := &ListNode{1, nil}

	head.Next = second
	second.Next = third

	sortList(head)

	cur := third
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Print("nil")
}
