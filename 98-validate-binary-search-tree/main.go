package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) (res bool) {
	var prev *TreeNode
	var check func(*TreeNode) bool
	check = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !check(root.Left) {
			return false
		}
		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev = root
		return check(root.Right)
	}
	return check(root)
}

func main() {

}
