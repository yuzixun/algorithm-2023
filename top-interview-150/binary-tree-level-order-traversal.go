package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	arr := []*TreeNode{root}

	for len(arr) > 0 {
		newArr := make([]*TreeNode, 0)
		levelArr := make([]int, 0)

		for _, item := range arr {
			levelArr = append(levelArr, item.Val)
			if item.Left != nil {
				newArr = append(newArr, item.Left)
			}
			if item.Right != nil {
				newArr = append(newArr, item.Right)
			}
		}
		result = append(result, levelArr)
		arr = newArr
	}
	return result
}
