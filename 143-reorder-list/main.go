package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 辅助数组
func reorderList1(head *ListNode) {
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	var guard ListNode
	p := &guard
	i, n, sign := 0, len(nodes), 1
	for n > 0 {
		p.Next = nodes[i]
		p = p.Next
		n--
		i = i + n*sign
		sign = -sign
	}
	p.Next = nil
}

// 不用额外空间: 找到中点, 倒转后半段, 合并
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	first := head
	for prev.Next != nil {
		first.Next, first = prev, first.Next
		prev.Next, prev = first, prev.Next
	}
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	reorderList(n1)
}
