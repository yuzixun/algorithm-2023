package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Right, root.Left)
}

func compare(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	return a.Val == b.Val && compare(a.Left, b.Right) && compare(a.Right, b.Left)
}
