package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLevel := treeLevel(root.Left)
	rightLevel := treeLevel(root.Right)

	if leftLevel == rightLevel {
		return countNodes(root.Right) + (1 << leftLevel)
	} else {
		return countNodes(root.Left) + (1 << rightLevel)
	}
}

func treeLevel(root *TreeNode) int {
	cur := root
	count := 0
	for cur != nil {
		cur = cur.Left
		count++
	}
	return count
}
