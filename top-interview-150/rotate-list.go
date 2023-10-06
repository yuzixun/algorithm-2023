package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	count := 0
	cur := head

	for cur != nil {
		count++
		cur = cur.Next
	}

	k %= count

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
