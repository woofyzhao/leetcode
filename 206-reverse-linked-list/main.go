package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}

func main() {

}
