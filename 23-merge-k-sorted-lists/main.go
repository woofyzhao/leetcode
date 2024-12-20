package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// An NodeHeap is a min-heap of ints.
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists1(lists []*ListNode) *ListNode {

	var h NodeHeap
	heap.Init(&h)
	for _, node := range lists {
		if node == nil {
			continue
		}
		heap.Push(&h, node)
	}

	var guard ListNode
	head := &guard
	for h.Len() > 0 {
		next := h[0]
		heap.Remove(&h, 0)
		println(next.Val)
		if next.Next != nil {
			heap.Push(&h, next.Next)
		}
		head.Next = next
		head = next
	}

	return guard.Next
}

// 直接归并排序更简单
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
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

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}
	mergeKLists([]*ListNode{list1, list2})
}
