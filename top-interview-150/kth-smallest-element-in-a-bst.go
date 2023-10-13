package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	i := 0

	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if i < k {
			i++
			arr = append(arr, node.Val)
		}
		dfs(node.Right)
	}
	dfs(root)
	return arr[k-1]
}
