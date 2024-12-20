package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	left := true
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, 0, n)
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if !left {
			slices.Reverse(tmp)
		}
		res = append(res, tmp)
		left = !left
		queue = queue[n:]
	}
	return
}
