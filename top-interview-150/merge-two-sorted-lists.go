package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := new(ListNode)
	head := node

	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			node.Next = p1
			p1 = p1.Next
		} else {
			node.Next = p2
			p2 = p2.Next
		}
		node = node.Next
	}

	if p1 != nil {
		node.Next = p1
	}
	if p2 != nil {
		node.Next = p2
	}

	return head.Next
}
