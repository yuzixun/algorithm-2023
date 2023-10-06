package main

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{
			Val: cur.Val,
		}
	}

	dummy := new(Node)
	p := dummy

	for cur := head; cur != nil; cur = cur.Next {
		p.Next = m[cur]
		p.Next.Next = m[cur.Next]
		p.Next.Random = m[cur.Random]
		p = p.Next
	}
	return dummy.Next
}
