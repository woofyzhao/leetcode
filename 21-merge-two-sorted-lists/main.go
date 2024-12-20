package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists0(list1 *ListNode, list2 *ListNode) *ListNode {
	guard := &ListNode{}
	p := guard
	for list1 != nil || list2 != nil {
		if list1 == nil {
			p.Next = list2
			break
		}
		if list2 == nil {
			p.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	return guard.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
