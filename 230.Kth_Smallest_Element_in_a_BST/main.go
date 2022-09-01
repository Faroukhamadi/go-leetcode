package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// nums := []int{}
	var nums []int
	inorder(root, nums, k)

	return nums[1]
}

func inorder(root *TreeNode, nums []int, k int) {
	if root == nil {
		return
	}

	inorder(root.Left, nums, k)
	nums[0]++
	if nums[0] == k {
		nums[1] = root.Val
		return
	}
	inorder(root.Right, nums, k)
}

func main() {

}
