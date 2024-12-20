package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}

func main() {
	
}
