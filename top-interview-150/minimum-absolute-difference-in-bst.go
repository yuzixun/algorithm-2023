package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt64, -1

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != -1 && root.Val-pre < res {
			res = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return res
}
