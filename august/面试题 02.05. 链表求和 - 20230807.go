package august

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	cur := new(ListNode)
	dummy := cur

	for l1 != nil && l2 != nil {
		cur.Next = new(ListNode)
		sum := l1.Val + l2.Val + carry
		carry, cur.Next.Val = sum/10, sum%10
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		cur.Next = new(ListNode)
		sum := l1.Val + carry
		carry, cur.Next.Val = sum/10, sum%10
		cur = cur.Next
		l1 = l1.Next
	}

	for l2 != nil {
		cur.Next = new(ListNode)
		sum := l2.Val + carry
		carry, cur.Next.Val = sum/10, sum%10
		cur = cur.Next
		l2 = l2.Next
	}

	if carry != 0 {
		cur.Next = new(ListNode)
		cur.Next.Val = carry
	}
	return dummy.Next
}
