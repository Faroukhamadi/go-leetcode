package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for slow.Next != nil {
		if fast.Next != nil {
			fast = fast.Next

			if fast.Next != nil {
				fast = fast.Next
			}
		}

		if slow == fast {
			return true
		}

		slow = slow.Next
	}

	return false
}

func main() {

}
