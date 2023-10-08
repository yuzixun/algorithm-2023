package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	arr := []*TreeNode{root}

	for len(arr) > 0 {
		var rightVal int
		newArr := make([]*TreeNode, 0)

		for _, item := range arr {
			rightVal = item.Val

			if item.Left != nil {
				newArr = append(newArr, item.Left)
			}
			if item.Right != nil {
				newArr = append(newArr, item.Right)
			}
		}

		res = append(res, rightVal)
		arr = newArr
	}

	return res
}
