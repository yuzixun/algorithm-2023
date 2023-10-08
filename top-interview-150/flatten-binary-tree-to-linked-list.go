package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	var pre *TreeNode
	var handle func(node *TreeNode)
	handle = func(node *TreeNode) {
		if node == nil {
			return
		}

		handle(node.Right)
		handle(node.Left)

		node.Right = pre
		node.Left = nil

		pre = node
	}
	handle(root)
}
