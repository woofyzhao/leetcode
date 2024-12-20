package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) (res *TreeNode) {
	if len(preorder) == 0 {
		return nil
	}
	res = &TreeNode{Val: preorder[0]}
	for i := range inorder {
		if inorder[i] == preorder[0] {
			res.Left = buildTree(preorder[1:i+1], inorder[0:i])
			res.Right = buildTree(preorder[i+1:], inorder[i+1:])
			return
		}
	}
	return
}

func main() {

}
