package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxIn(nums []int) int {
	res := nums[0]
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}

const MinVal int = -10000

func maxPathSum(root *TreeNode) int {
	maxSum = MinVal
	maxPath(root)
	return maxSum
}

func maxPath(root *TreeNode) int {
	if root == nil {
		return MinVal
	}
	maxLeft := maxPath(root.Left)
	maxRight := maxPath(root.Right)
	max := maxIn([]int{root.Val, maxLeft + root.Val, maxRight + root.Val})
	if max > maxSum {
		maxSum = max
	}
	if maxLeft+maxRight+root.Val > maxSum {
		maxSum = maxLeft + maxRight + root.Val
	}
	return max
}

func main() {
	root := &TreeNode{Val: -3}
	fmt.Println(maxPathSum(root))
}
