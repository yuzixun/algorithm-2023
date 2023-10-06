package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{0, head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		if prev.Next.Val == prev.Next.Next.Val {
			cur := prev.Next
			x := cur.Val
			for cur.Next != nil && x == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
	}

	return dummy.Next
}
