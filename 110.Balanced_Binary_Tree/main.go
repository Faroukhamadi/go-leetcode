package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var h = helper(root)
	if h == -1 {
		return false
	}

	return true
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var lh = helper(root.Left)

	if lh == -1 {
		return -1
	}

	var rh = helper(root.Right)

	if rh == -1 {
		return -1
	}

	if math.Abs(float64(lh)-float64(rh)) > 1 {
		return -1
	}

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func main() {
}
