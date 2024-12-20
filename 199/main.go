package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	right := rightSideView(root.Right)
	res = append(res, right...)

	left := rightSideView(root.Left)
	if len(left) > len(right) {
		res = append(res, left[len(right):]...)
	}
	return
}
