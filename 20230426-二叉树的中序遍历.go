package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal2(root *TreeNode) []int {
	var (
		res []int
		a   func(node *TreeNode)
	)
	a = func(node *TreeNode) {
		if node == nil {
			return
		}
		a(node.Left)
		res = append(res, node.Val)
		a(node.Right)
	}
	a(root)
	return res
}
