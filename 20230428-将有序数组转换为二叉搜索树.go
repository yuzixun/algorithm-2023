package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	node := new(TreeNode)
	index := (len(nums) - 1) / 2
	node.Val = nums[index]
	node.Left = sortedArrayToBST(nums[:index])
	node.Right = sortedArrayToBST(nums[index+1:])
	return node
}
