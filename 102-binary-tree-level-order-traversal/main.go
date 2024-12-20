package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Elem struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []Elem{{Node: root}}
	lastLevel := -1
	for len(queue) > 0 {
		head := queue[0]
		if head.Level != lastLevel {
			res = append(res, []int{})
			lastLevel = head.Level
		}
		res[len(res)-1] = append(res[len(res)-1], head.Node.Val)
		if head.Node.Left != nil {
			queue = append(queue, Elem{Node: head.Node.Left, Level: head.Level + 1})
		}
		if head.Node.Right != nil {
			queue = append(queue, Elem{Node: head.Node.Right, Level: head.Level + 1})
		}
		queue = queue[1:]
	}
	return
}

func main() {
	root := TreeNode{Val: 100}
	fmt.Println(levelOrder(&root))
}
