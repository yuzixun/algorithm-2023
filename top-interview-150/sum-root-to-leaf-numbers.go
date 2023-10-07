package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return sumNode(root, 0)
}

func sumNode(node *TreeNode, res int) int {
	if node == nil {
		return 0
	}
	tmp := res*10 + node.Val
	if node.Right == nil && node.Left == nil {
		return tmp
	}
	return sumNode(node.Left, tmp) + sumNode(node.Right, tmp)
}
