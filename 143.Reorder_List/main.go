package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

}

func splitInTwoLists(node *ListNode) *ListNode {
	middle := middleNode(node)
	headList2 := middle.Next
	middle.Next = nil
	return headList2
}

func middleNode(node *ListNode) *ListNode {
	slow, fast := node, node
	for slow != nil && fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func reverseList(node *ListNode) *ListNode {
	var previous, current *ListNode = nil, node
	for current != nil {
		previous, current, current.Next = current, current.Next, previous
	}
	return previous
}

func mergeTwoLists(head1, head2 *ListNode) {
	var current1, current2 *ListNode = head1, head2
	for current1 != nil && current2 != nil {
		current1, current2, current1.Next, current2.Next = current1.Next, current2.Next, current2, current1.Next
	}
}

func main() {
}
