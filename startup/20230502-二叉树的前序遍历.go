package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var (
		res    []int
		helper func(root *TreeNode)
	)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return res
}
