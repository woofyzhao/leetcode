package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (res int) {
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := height(root.Left)
		right := height(root.Right)
		if left+right > res {
			res = left + right
		}
		return max(left, right) + 1
	}
	height(root)
	return
}
