package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) (res bool) {
	res = true
	var checkHeight func(r *TreeNode) int
	checkHeight = func(r *TreeNode) int {
		if !res || r == nil {
			return 0
		}
		left := checkHeight(r.Left)
		right := checkHeight(r.Right)
		if left-right > 1 || right-left > 1 {
			res = false
		}
		return max(left, right) + 1
	}
	checkHeight(root)
	return
}

func main() {

}
