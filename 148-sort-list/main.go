package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	left, right := head, slow.Next
	slow.Next = nil
	return merge(sortList(left), sortList(right))
}

func merge(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val > b.Val {
		return merge(b, a)
	}
	a.Next = merge(a.Next, b)
	return a
}

func main() {
	n1 := &ListNode{Val: 5}
	n2 := &ListNode{Val: 9}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n4 = sortList(n1)
	n1.Val = 5
}
