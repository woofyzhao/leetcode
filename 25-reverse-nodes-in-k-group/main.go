package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	var prev *ListNode = nil
	for first != last {
		first.Next, first, prev = prev, first.Next, first
	}
	return prev
}

func main() {

}
