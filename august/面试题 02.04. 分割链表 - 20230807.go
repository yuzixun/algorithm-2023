package august

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	lDummy, rDummy := &ListNode{}, &ListNode{}
	lHelper, rHelper := lDummy, rDummy
	cur := head

	for cur != nil {
		if cur.Val < x {
			lDummy.Next = cur
			lDummy = lDummy.Next
		} else {
			rDummy.Next = cur
			rDummy = rDummy.Next
		}
		cur = cur.Next
	}

	rDummy.Next = nil
	lDummy.Next = rHelper.Next
	return lHelper.Next
}
