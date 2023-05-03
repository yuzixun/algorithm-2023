package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var isSymmetricLeftRight func(left *TreeNode, right *TreeNode) bool

	isSymmetricLeftRight = func(left *TreeNode, right *TreeNode) bool {
		if (left == nil && right != nil) || (right == nil && left != nil) {
			return false
		}
		if left == nil && right == nil {
			return true
		}
		if left.Val != right.Val {
			return false
		}

		return isSymmetricLeftRight(left.Left, right.Right) && isSymmetricLeftRight(left.Right, right.Left)
	}

	return isSymmetricLeftRight(root.Left, root.Right)
}
