package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var add func(a *ListNode, b *ListNode, carry int) *ListNode
	add = func(a *ListNode, b *ListNode, carry int) *ListNode {
		if a == nil {
			if b == nil {
				if carry > 0 {
					return &ListNode{Val: carry}
				}
				return nil
			}
			return add(b, a, carry)
		}
		var bv int
		var bp *ListNode
		if b != nil {
			bv = b.Val
			bp = b.Next
		}
		newCarry := (a.Val + bv + carry) / 10
		a.Val = (a.Val + bv + carry) % 10
		a.Next = add(a.Next, bp, newCarry)
		return a
	}
	return add(l1, l2, 0)
}

func main() {

}
