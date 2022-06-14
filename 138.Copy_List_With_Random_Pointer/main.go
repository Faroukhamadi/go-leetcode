package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var mp = make(map[*Node]*Node)
	var cur = head
	var clone *Node = nil

	// Create a new list without links
	for cur != nil {
		clone = &Node{cur.Val, nil, nil}
		mp[cur] = clone
		cur = cur.Next
	}

	cur = head
	// Traverse list and link the clone node
	for cur != nil {
		clone = mp[cur]
		clone.Next = mp[cur.Next]
		clone.Random = mp[cur.Random]

		cur = cur.Next
	}

	return mp[head]
}

func main() {

}
