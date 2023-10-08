package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)

	arr := []*TreeNode{root}
	for len(arr) > 0 {
		count, sum := 0, 0
		newArr := make([]*TreeNode, 0)

		for _, item := range arr {
			count++
			sum += item.Val
			if item.Left != nil {
				newArr = append(newArr, item.Left)
			}
			if item.Right != nil {
				newArr = append(newArr, item.Right)
			}
		}

		result = append(result, float64(sum)/float64(count))
		arr = newArr
	}

	return result
}
