package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	front, back := head, head
	for i := 0; front != nil && i < n; i++ {
		front = front.Next
	}
	if front == nil {
		return head.Next
	}
	for front.Next != nil {
		front = front.Next
		back = back.Next
	}
	back.Next = back.Next.Next
	return head
}

func main() {

}
