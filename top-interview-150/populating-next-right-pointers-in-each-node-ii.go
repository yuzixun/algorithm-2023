package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummy := new(Node)
		pre := dummy
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}

//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//
//	arr := []*Node{root}
//	for len(arr) > 0 {
//		var (
//			pre    *Node
//			cur    *Node
//			newArr []*Node
//		)
//
//		for _, node := range arr {
//			cur = node.Left
//			if cur != nil {
//				newArr = append(newArr, node.Left)
//				if pre != nil {
//					pre.Next = cur
//				}
//				pre = cur
//			}
//
//			cur = node.Right
//			if cur != nil {
//				newArr = append(newArr, node.Right)
//				if pre != nil {
//					pre.Next = cur
//				}
//				pre = cur
//			}
//
//		}
//
//		arr = newArr
//	}
//
//	return root
//}
