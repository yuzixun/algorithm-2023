package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := inorder[len(inorder)-1]
	if len(inorder) == 1 {
		return &TreeNode{Val: val}
	}

	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			index = i
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}
