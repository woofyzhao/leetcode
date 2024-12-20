package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// collect to a stack and pop common nodes
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	collect := func(head *ListNode) []*ListNode {
		var s []*ListNode
		for head != nil {
			s = append(s, head)
			head = head.Next
		}
		return s
	}

	s1 := collect(headA)
	s2 := collect(headB)

	i, j := len(s1)-1, len(s2)-1
	if s1[i] != s2[j] {
		return nil
	}

	for i >= 0 && j >= 0 && s1[i] == s2[j] {
		i--
		j--
	}
	return s1[i+1]
}

// calc len diff and adjust heads to be identical distance to the common node
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	length := func(head *ListNode) (res int) {
		for head != nil {
			res += 1
			head = head.Next
		}
		return
	}
	lenA := length(headA)
	lenB := length(headB)
	if lenA < lenB {
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}
	for i := 0; i < lenA-lenB; i++ {
		headA = headA.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

// 拼接遍历 (距交点距离一致)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}

func main() {

}
