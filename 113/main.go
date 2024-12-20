package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	var dfs func(*TreeNode, []int, int)
	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		sum += node.Val

		if sum == targetSum && node.Left == nil && node.Right == nil {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
		}

		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}

	dfs(root, []int(nil), 0)

	return res
}

func main() {

}
