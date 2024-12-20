package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > 1 {
		head.Next = reverseBetween(head.Next, left-1, right-1)
		return head
	}
	var prev *ListNode
	newTail := head
	for i := 0; i < right; i++ {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	newTail.Next = head
	return prev
}
