package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	var (
		pre *TreeNode
		res = true
		dfs func(root *TreeNode)
	)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil {
			res = res && root.Val > pre.Val
		}
		pre = root
		dfs(root.Right)
	}

	dfs(root)
	return res
}
