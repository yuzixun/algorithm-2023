package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var (
		res    []int
		helper func(root *TreeNode)
	)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)
		helper(root.Right)
		res = append(res, root.Val)
	}
	helper(root)
	return res
}
