package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	token := 0

	for l1 != nil && l2 != nil {
		cur := token + l1.Val + l2.Val
		p.Next, token = &ListNode{Val: cur % 10}, cur/10

		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		cur := token + l1.Val
		p.Next, token = &ListNode{Val: cur % 10}, cur/10

		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		cur := token + l2.Val
		p.Next, token = &ListNode{Val: cur % 10}, cur/10

		p = p.Next
		l2 = l2.Next
	}
	if token != 0 {
		p.Next = &ListNode{Val: token}
	}
	return dummy.Next
}
