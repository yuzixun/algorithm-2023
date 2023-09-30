package august

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {

	pre, cur := new(ListNode), head
	dataMap := map[int]struct{}{}

	for ; cur != nil; cur = cur.Next {
		_, exists := dataMap[cur.Val]
		if exists {
			pre.Next = cur.Next
		} else {
			dataMap[cur.Val] = struct{}{}
			pre = cur
		}
	}

	return head
}
