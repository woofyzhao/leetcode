package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Val != head.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
	p := head
	for p != nil && p.Val == head.Val {
		p = p.Next
	}
	return deleteDuplicates(p)
}

func main() {

}
