package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) (res int) {
	res = 1
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left == nil {
		return widthOfBinaryTree(root.Right)
	}
	if root.Right == nil {
		return widthOfBinaryTree(root.Left)
	}
	m := make(map[int]int) // level -> tag
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level int, tag int) {
		if node == nil {
			return
		}
		if m[level] > 0 {
			res = max(res, tag-m[level]+1)
		} else {
			m[level] = tag
		}
		dfs(node.Left, level+1, tag<<1)
		dfs(node.Right, level+1, (tag<<1)+1)
	}
	dfs(root, 0, 1)
	return res
}
