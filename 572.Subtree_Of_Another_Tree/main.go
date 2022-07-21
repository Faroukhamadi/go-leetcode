package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return isEqual(root.Left, subRoot.Left) && isEqual(root.Right, subRoot.Right)
}
