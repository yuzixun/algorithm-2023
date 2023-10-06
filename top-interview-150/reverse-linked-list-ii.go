package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}

	p := dummy
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	cur := p.Next
	var pre *ListNode
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	p.Next.Next = cur
	p.Next = pre

	return dummy.Next
}

//func reverseBetween(head *ListNode, left int, right int) *ListNode {
//	if left == 1 {
//		return reverseN(head, right)
//	}
//
//	head.Next = reverseBetween(head.Next, left-1, right-1)
//	return head
//}
//
//var successor *ListNode
//
//func reverseN(head *ListNode, n int) *ListNode {
//	if n == 1 {
//		successor = head.Next
//		return head
//	}
//
//	last := reverseN(head.Next, n-1)
//	head.Next.Next = head
//	head.Next = successor
//	return last
//}
