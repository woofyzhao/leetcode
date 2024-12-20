package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev, head = prev.Next, head.Next
	}
	return true
}

func main() {

}
