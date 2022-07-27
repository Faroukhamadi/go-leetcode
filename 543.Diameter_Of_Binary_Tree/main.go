package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, mx := getHtMx(root)
	return max(mx-1, 0)
}

func getHtMx(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lht, lmx := getHtMx(node.Left)
	rht, rmx := getHtMx(node.Right)
	mx := max(lmx, rmx)
	ht := 1 + max(lht, rht)
	mx = max(mx, 1+lht+rht)
	return ht, mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}
