package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 转换为求共同交点问题
func detectCycle1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	p := head
	for p != fast {
		p = p.Next
		if fast.Next == slow {
			fast = head
		} else {
			fast = fast.Next
		}
	}
	return p
}

// 利用2倍关系  计算可得一定在某处对齐
// 2 * (a + b) = a + b + nc
// a = nc - b
// a = (n-1-2sum)c + (c-b) 即 相差 n圈
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			for head != slow {
				head, slow = head.Next, slow.Next
			}
			return head
		}
	}
	return nil
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	detectCycle(n1)
}
